package f_init

import (
	"context"

	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/f_global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := f_global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOGGER.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.LOGGER.Info("redis connect ping response:", zap.String("pong", pong))
		f_global.GVA_REDIS = client
	}
}
