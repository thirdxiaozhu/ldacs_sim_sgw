package f_init

import (
	_ "ldacs_sim_sgw/pkg/forward_module/source/example"
	_ "ldacs_sim_sgw/pkg/forward_module/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
