package main

import (
	"ldacs_sim_sgw/internal/core"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/initialize"
	backward "ldacs_sim_sgw/pkg/backward_module"
	forward "ldacs_sim_sgw/pkg/forward_module"
	ldacscore "ldacs_sim_sgw/pkg/ldacs_core"
	"os"
	"os/signal"
	"syscall"
)

// #cgo CFLAGS: -I /usr/local/include/ldacs
// #cgo LDFLAGS:  -lldacscore -lldacsmsgcore  -lldacsnetcore -lldacsutilcore -lldacsrolecore -lgmssl -lm -lyaml -lcjson -lbase64
// #include <ldacs_core/ldacs_core.h>
import "C"

func systemTeardown() {
}

func mainloop() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-exitSignal

	systemTeardown()
}

func main() {

	//para := make([]uint8, 4)
	//C.generate_rand((*C.uchar)(unsafe.Pointer(&para[0])))
	//fmt.Println(para, binary.BigEndian.Uint32(para))

	global.VP = core.InitViper()   // 初始化Viper
	global.LOGGER = core.InitZap() // 初始化zap日志库
	global.DB = initialize.Gorm()  // gorm连接数据库
	initialize.DBList()
	if global.DB != nil {
		//initialize.RegisterTables() // 初始化表
		ldacscore.InitCoreModule()
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	/* run backward module */
	go backward.ListenAndServe(":7777", ldacscore.MakeLdacsHandler())
	/* run forward module */
	go forward.RunForward()

	/* waiting for all go routine exit */
	mainloop()
}
