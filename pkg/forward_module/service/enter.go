package service

import (
	"ldacs_sim_sgw/pkg/forward_module/service/example"
	"ldacs_sim_sgw/pkg/forward_module/service/system"
	"ldacs_sim_sgw/pkg/ldacs_core/service"
)

type ServiceGroup struct {
	SystemServiceGroup            system.ServiceGroup
	ExampleServiceGroup           example.ServiceGroup
	Ldacs_sgw_forwardServiceGroup service.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
