package ldacscore

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/handle"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"os"
)

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		model.AccountPlane{},
		model.AccountFlight{},
		model.AccountAuthz{},
		model.AccountAs{},
		model.AccountGs{},
		model.AccountGsc{},
		model.AuditAsRaw{},
		model.AuthzPlane{},
		model.AuthcState{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
func InitCoreModule() {
	RegisterTables()
}
func MakeLdacsHandler() *handle.LdacsHandler {
	return &handle.LdacsHandler{}
}
