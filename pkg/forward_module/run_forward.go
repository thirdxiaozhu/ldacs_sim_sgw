package forward_module

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/f_core"
	"ldacs_sim_sgw/pkg/forward_module/f_global"
	"ldacs_sim_sgw/pkg/forward_module/f_init"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// RunForward @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.5.9
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func RunForward() {
	f_global.GVA_VP = f_core.ForwardViper() // 初始化Viper
	f_init.OtherInit()
	zap.ReplaceGlobals(global.LOGGER)
	f_init.Timer()
	//global.DB = initialize.Gorm() // gorm连接数据库
	//initialize.DBList()
	if global.DB != nil {
		f_init.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		//db, _ := global.DB.DB()
		//defer db.Close()
	}
	f_core.RunWindowsServer()
}
