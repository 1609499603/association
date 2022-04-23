package initialize

import (
	"association/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Redis 初始化redis
func Redis() {
	redisCfg := global.ASS_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})
	result, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.ASS_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.ASS_LOG.Info("redis connect ping response:", zap.String("result", result))
		global.ASS_REDIS = client
	}
}
