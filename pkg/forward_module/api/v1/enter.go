package v1

import (
	"ldacs_sim_sgw/pkg/forward_module/api/v1/example"
	"ldacs_sim_sgw/pkg/forward_module/api/v1/ldacs_sgw_forward"
	"ldacs_sim_sgw/pkg/forward_module/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup            system.ApiGroup
	ExampleApiGroup           example.ApiGroup
	Ldacs_sgw_forwardApiGroup ldacs_sgw_forward.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
