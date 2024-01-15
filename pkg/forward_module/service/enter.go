package service

import (
	"ldacs_sim_sgw/pkg/forward_module/service/example"
	"ldacs_sim_sgw/pkg/forward_module/service/ldacs_sgw_forward"
	"ldacs_sim_sgw/pkg/forward_module/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup            system.ServiceGroup
	ExampleServiceGroup           example.ServiceGroup
	Ldacs_sgw_forwardServiceGroup ldacs_sgw_forward.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
