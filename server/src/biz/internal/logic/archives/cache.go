package archives

import (
	"biz/internal/svc"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const cacheKey string = "cache:archives:page:"
const pageSize int64 = 10

func GetArchivesBlogIds(ctx context.Context, svc svc.ServiceContext, page int64) ([]string, error) {
	var latestBlogIds []string
	ckey := fmt.Sprintf("%s%d", cacheKey, page)

	cacheData, err := svc.Rds.GetCtx(ctx, ckey)

	if err == redis.Nil || len(cacheData) == 0 { // Redis 缓存中不存在，查询 MongoDB
		// MongoDB 查询选项
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{Key: "updateAt", Value: -1}})
		findOptions.SetLimit(10)
		findOptions.SetSkip(page * 10) // 根据页码跳过前面的记录

		// 设置投影，只返回 _id 字段
		projection := bson.D{{Key: "_id", Value: 1}}
		findOptions.SetProjection(projection)

		conn := svc.BlogModel.GetConn()
		cursor, err := conn.Collection.Find(context.Background(), bson.D{}, findOptions)
		if err != nil {
			logx.Error("MongoDB query error:", err)
			return nil, err
		}
		defer cursor.Close(context.Background())

		// 读取 MongoDB 返回的结果，只保存ID
		for cursor.Next(context.TODO()) {
			var record struct {
				ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
			}
			if err = cursor.Decode(&record); err != nil {
				logx.Error("MongoDB decode error:", err)
				return nil, err
			}
			latestBlogIds = append(latestBlogIds, record.ID.Hex())
		}

		if err := cursor.Err(); err != nil {
			log.Println("Cursor error:", err)
			return nil, err
		}

		// 将ID列表存入 Redis 缓存，设置过期时间（如 1 小时）
		dataToCache, err := json.Marshal(latestBlogIds)
		if err != nil {
			log.Println("Failed to marshal data for Redis:", err)
			return nil, err
		}
		go func(ckey string) {
			saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			sub_err := svc.Rds.SetexCtx(saveCtx, ckey, string(dataToCache), 3600)
			if sub_err != nil {
				logx.Error("Failed to set Redis cache:", sub_err)
			}
		}(ckey)

	} else if err != nil {
		logx.Error("Redis error:", err)
		return nil, err
	} else {
		err = json.Unmarshal([]byte(cacheData), &latestBlogIds)
		if err != nil {
			logx.Info("Failed to unmarshal cached data:", err)
			go clearCache(svc.Rds)
			return nil, err
		}
	}
	return latestBlogIds, nil
}

// 清除分页缓存
func clearCache(rdb *redis.Redis) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	pattern := fmt.Sprintf("%s*", cacheKey) // 匹配所有分页缓存键
	var cursor uint64 = 0
	for {
		keys, nextCursor, err := rdb.ScanCtx(ctx, cursor, pattern, 10)
		if err != nil {
			logx.Errorf("Failed to scan keys: %v", err)
			panic(err)
		}

		// 打印获取到的键
		if len(keys) > 0 {
			_, err = rdb.DelCtx(ctx, keys...)
			if err != nil {
				logx.Error("Redis delete error:", err)
				panic(err)
			}
		}

		// 如果游标返回为 0，表示扫描结束
		if nextCursor == 0 {
			break
		}

		// 更新游标，继续扫描
		cursor = nextCursor
	}
	// :max也一起删了
}

// 获得页数(博文数量 / 10)
//
// 页码从0开始依次递增,
// 如果总量为5,那么将返回1页;总量为15,返回2页
func GetArchivePageNumber(ctx context.Context, conn *monc.Model, rds *redis.Redis) (int64, error) {
	var totalPages int64
	ckey := fmt.Sprintf("%smax", cacheKey)

	cacheData, err := rds.GetCtx(ctx, ckey)

	if err == redis.Nil || cacheData == "" {
		return initStatPageNum(ctx, conn, rds)
	} else if err != nil {
		return 0, err
	} else {
		err = json.Unmarshal([]byte(cacheData), &totalPages)
		if err != nil {
			logx.Info("Failed to unmarshal cached data:", err)
			go clearCache(rds)
			return 0, err
		}
		return totalPages, nil
	}
}

// Initialize to statistic page numbers
// and store to redis
func initStatPageNum(ctx context.Context, conn *monc.Model, rds *redis.Redis) (int64, error) {
	var totalPages int64
	count, err := conn.Collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}
	// 计算总页数
	totalPages = (count + pageSize - 1) / pageSize // 向上取整
	dataToCache, err := json.Marshal(totalPages)

	go func(ckey string) {
		saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = rds.SetexCtx(saveCtx, ckey, string(dataToCache), 3600)
		if err != nil {
			logx.Error("Failed to set Redis cache:", err)
		}
	}(fmt.Sprintf("%smax", cacheKey))

	return totalPages, nil
}
