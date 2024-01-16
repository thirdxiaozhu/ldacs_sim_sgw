package response

import "ldacs_sim_sgw/pkg/forward_module/f_config"

type SysConfigResponse struct {
	Config f_config.ServerConfig `json:"config"`
}
