package response

import "ldacs_sim_sgw/pkg/forward_module/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
