package initialize

import (
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
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
		ldacs_sgw_forward.AccountPlane{},
		ldacs_sgw_forward.AccountFlight{},
		ldacs_sgw_forward.AccountAuthz{},
		ldacs_sgw_forward.AuthzPlane{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
