package svc

import (
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func GetRedis(c []cache.NodeConf) (*redis.Redis, error) {
	if len(c) == 0 || cache.TotalWeights(c) <= 0 {
		err := errors.New("no cache nodes in cache")
		logx.Alert(err.Error())
		return nil, err
	}
	return redis.MustNewRedis(c[0].RedisConf), nil
}
