package model

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"os"
)

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		AccountPlane{},
		AccountFlight{},
		AccountAuthz{},
		AccountAs{},
		AccountGs{},
		AccountGsc{},
		AuditAsRaw{},
		AuthzPlane{},
		AuthcState{},
		State{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
