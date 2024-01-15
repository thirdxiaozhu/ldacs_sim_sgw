package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"ldacs_sim_sgw/pkg/forward_module/forward_global"
	"ldacs_sim_sgw/pkg/forward_module/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(forward_global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(forward_global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	forward_global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
