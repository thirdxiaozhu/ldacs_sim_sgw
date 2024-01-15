package main

import (
	"fmt"
	backward "ldacs_sim_sgw/pkg/backward_module"
	forward "ldacs_sim_sgw/pkg/forward_module"
	ldacscore "ldacs_sim_sgw/pkg/ldacs_core"
	"os"
	"os/signal"
	"syscall"
)

// #cgo CFLAGS: -I /usr/local/include/ldacs
// #cgo LDFLAGS:  -lldacscore -lldacsmsgcore  -lldacsnetcore -lldacsutilcore -lldacsrolecore -lgmssl -lm -lyaml -lcjson
// #include <ldacs_core/ldacs_core.h>
import "C"

func systemTeardown() {
	fmt.Println("!!!!!!!!!!!!!!!")
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

	/* run backward module */
	go backward.ListenAndServe(":7777", ldacscore.MakeLdacsHandler())
	/* run forward module */
	go forward.RunForward()

	/* waiting for all go routine exit */
	mainloop()
}
