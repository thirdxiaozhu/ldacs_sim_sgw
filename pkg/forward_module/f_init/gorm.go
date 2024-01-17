package f_init

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/model/example"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
	"os"
)

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysExportTemplate{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, ldacs_sgw_forward.AirStation{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
