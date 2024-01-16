package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	internal "ldacs_sim_sgw/internal/core/inside"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"os"
)

// InitZap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func InitZap() (logger *zap.Logger) {
	if ok, _ := util.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
