package initialize

import (
	"context"

	"ldacs_sim_sgw/pkg/forward_module/forward_global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := forward_global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		forward_global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		forward_global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		forward_global.GVA_REDIS = client
	}
}
