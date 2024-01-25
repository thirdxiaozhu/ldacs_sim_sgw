package f_core

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/f_global"

	"ldacs_sim_sgw/pkg/forward_module/f_init"
	"ldacs_sim_sgw/pkg/forward_module/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if f_global.GVA_CONFIG.System.UseMultipoint || f_global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		f_init.Redis()
	}
	if f_global.GVA_CONFIG.System.UseMongo {
		err := f_init.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.DB != nil {
		system.LoadAll()
	}

	Router := f_init.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONFIG.System.ForwardAddr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Info("server run success on ", zap.String("address", address))
	global.LOGGER.Error(s.ListenAndServe().Error())
}
