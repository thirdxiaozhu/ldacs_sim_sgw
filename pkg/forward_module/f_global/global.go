package f_global

import (
	"github.com/qiniu/qmgo"
	"ldacs_sim_sgw/pkg/forward_module/f_config"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"ldacs_sim_sgw/pkg/forward_module/utils/timer"

	"golang.org/x/sync/singleflight"

	"github.com/redis/go-redis/v9"

	"github.com/spf13/viper"
)

var (
	GVA_REDIS               *redis.Client
	GVA_MONGO               *qmgo.QmgoClient
	GVA_CONFIG              f_config.ServerConfig
	GVA_VP                  *viper.Viper
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)
