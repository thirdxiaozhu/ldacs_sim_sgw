package main

import (
	"github.com/hdt3213/godis/lib/logger"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/core"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/initialize"
	backward "ldacs_sim_sgw/pkg/backward_module"
	forward "ldacs_sim_sgw/pkg/forward_module"
	ldacscore "ldacs_sim_sgw/pkg/ldacs_core"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

//// #cgo CFLAGS: -I /usr/local/include/ldacs
//// #cgo LDFLAGS:  -lldacscore -lldacsmsgcore  -lldacsnetcore -lldacsutilcore -lldacsrolecore -lgmssl -lm -lyaml -lcjson -lbase64
//// #include <ldacs_core/ldacs_core.h>
//import "C"

func systemTeardown() {
}

func mainloop() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-exitSignal

	systemTeardown()
}

func main() {

	global.VP = core.InitViper()   // 初始化Viper
	global.LOGGER = core.InitZap() // 初始化zap日志库
	global.DB = initialize.Gorm()  // gorm连接数据库
	// wsl访问sqlite设置jdbc URL:  jdbc:sqlite:file:\\wsl$\Ubuntu\home\jiaxv\go\src\ldacs_sim_sgw\resources\ld_sql.db?nolock=1
	global.KeyDB = initialize.KeyGorm()

	initialize.DBList()
	if global.DB != nil && global.KeyDB != nil {
		//initialize.RegisterTables() // 初始化表
		ldacscore.InitCoreModule()
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	} else {
		global.LOGGER.Error("Fatal:", zap.String("Reason", "The database has not initialize correctly."))
		return
	}

	logger.Warn(strconv.FormatUint(uint64(global.CONFIG.System.SgwUA), 10))

	/* run backward module */
	go backward.ListenAndServe(":6666", ldacscore.MakeLdacsHandler())
	/* run forward module */
	go forward.RunForward()

	/* waiting for all go routine exit */
	mainloop()
}
