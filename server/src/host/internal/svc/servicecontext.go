package svc

import (
	"context"
	model "host/db/mongo"
	"host/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	*MongoModels
}

type MongoModels struct {
	BlogModel model.BlogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 启用了auth，因此需要先注册，redis中写入app与token
	// https://juejin.cn/post/7044185614811398174
	rds := redis.MustNewRedis(c.Redis.RedisConf)
	ctx := context.Background()

	var err error
	exists, err := rds.ExistsCtx(ctx, c.Redis.Key)
	if err != nil {
		logx.Errorf("can not judge whether host's rpc key is existed: %v", err)
		panic(err)
	}

	if exists {
		_, err = rds.DelCtx(ctx, c.Redis.Key)
		if err != nil {
			logx.Errorf("can not delete existed host's rpc key: %v", err)
			panic(err)
		}
	}

	for _, v := range c.AuthApps {
		err = rds.HsetCtx(ctx, c.Redis.Key, v.App, v.Token)
		if err != nil {
			logx.Infof("unable to create hset [%s] in redis: %v", v.App, err)
			continue
		}
	}

	// 建立数据库链接
	bm := model.NewBlogModelByCfg(c.MongoConfig.ToUri(), c.MongoConfig, c.MongoConfig.NodeConfs)

	return &ServiceContext{
		Config: c,
		MongoModels: &MongoModels{
			BlogModel: bm,
		},
	}
}
