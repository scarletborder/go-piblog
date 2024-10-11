package recommend

import (
	"context"
	"log"

	model "biz/db/mongo"
	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetLatestBlogIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLatestBlogIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestBlogIdsLogic {
	return &GetLatestBlogIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLatestBlogIdsLogic) GetLatestBlogIds() (resp *types.LatestBlogIdsResp, err error) {
	var latestRecords []string
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updateAt", Value: -1}})
	findOptions.SetLimit(5)

	// 设置投影，只返回 _id 字段
	projection := bson.D{{Key: "_id", Value: 1}}
	findOptions.SetProjection(projection)

	conn := l.svcCtx.BlogModel.GetConn()
	cursor, err := conn.Collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		logx.Error(err)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.TODO()) {
		var record model.Blog
		if err = cursor.Decode(&record); err != nil {
			logx.Error(err)
			return
		}
		latestRecords = append(latestRecords, record.ID.Hex())
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	resp = &types.LatestBlogIdsResp{IDs: latestRecords}
	return
}
