package b_init

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
	"os"
)

func RegisterBackwardTables() {
	db := global.DB
	err := db.AutoMigrate(
		system.SysApi{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
