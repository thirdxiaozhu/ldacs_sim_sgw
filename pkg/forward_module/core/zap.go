package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	internal "ldacs_sim_sgw/pkg/forward_module/core/inside"
	"ldacs_sim_sgw/pkg/forward_module/forward_global"
	"ldacs_sim_sgw/pkg/forward_module/utils"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(forward_global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", forward_global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(forward_global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if forward_global.GVA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
