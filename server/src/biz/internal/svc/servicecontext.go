package svc

import (
	model "biz/db/mongo"
	"biz/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Rds    *redis.Redis
	*MongoModels
}

type MongoModels struct {
	BlogModel model.BlogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 建立数据库链接
	bm := model.NewBlogModelByCfg(c.MongoConfig.ToUri(), c.MongoConfig, c.MongoConfig.NodeConfs)
	rds, err := GetRedis(c.MongoConfig.NodeConfs)
	if err != nil {
		panic(err)
	}
	// a := cache.ClusterConf{}

	return &ServiceContext{
		Config: c,
		Rds:    rds,
		MongoModels: &MongoModels{
			BlogModel: bm,
		},
	}
}
