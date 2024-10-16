package logic

import (
	"context"
	"fmt"
	"host/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const cacheKey string = "cache:archives:page:"

func clearCache(c config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rds := redis.MustNewRedis(c.Redis.RedisConf)

	pattern := fmt.Sprintf("%s*", cacheKey) // 匹配所有分页缓存键
	var cursor uint64 = 0
	for {
		keys, nextCursor, err := rds.ScanCtx(ctx, cursor, pattern, 10)
		if err != nil {
			logx.Errorf("Failed to scan keys: %v", err)
			panic(err)
		}

		// 打印获取到的键
		if len(keys) > 0 {
			_, err = rds.DelCtx(ctx, keys...)
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
