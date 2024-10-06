package svc

import (
	model "biz/db/mongo"
	"biz/internal/config"
)

type ServiceContext struct {
	Config config.Config
	*MongoModels
}

type MongoModels struct {
	BlogModel model.BlogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 建立数据库链接
	bm := model.NewBlogModelByCfg(c.MongoConfig.ToUri(), c.MongoConfig, c.MongoConfig.NodeConfs)

	return &ServiceContext{
		Config: c,
		MongoModels: &MongoModels{
			BlogModel: bm,
		},
	}
}
