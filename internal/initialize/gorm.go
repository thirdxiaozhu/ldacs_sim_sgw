package initialize

import (
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"ldacs_sim_sgw/internal/global"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

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
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
