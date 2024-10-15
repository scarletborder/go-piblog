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
	keys, err := rds.KeysCtx(ctx, fmt.Sprintf("%s*", cacheKey))
	if err != nil {
		logx.Error("Redis keys error:", err)
		return
	}

	// 删除所有匹配的分页缓存键
	if len(keys) > 0 {
		_, err = rds.DelCtx(ctx, keys...)
		if err != nil {
			logx.Error("Redis delete error:", err)
		}
	}
}
