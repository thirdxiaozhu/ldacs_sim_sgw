package router

import (
	"ldacs_sim_sgw/pkg/forward_module/router/example"
	"ldacs_sim_sgw/pkg/forward_module/router/ldacs_sgw_forward"
	"ldacs_sim_sgw/pkg/forward_module/router/system"
)

type RouterGroup struct {
	System            system.RouterGroup
	Example           example.RouterGroup
	Ldacs_sgw_forward ldacs_sgw_forward.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
