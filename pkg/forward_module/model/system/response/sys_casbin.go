package response

import (
	"ldacs_sim_sgw/pkg/forward_module/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
