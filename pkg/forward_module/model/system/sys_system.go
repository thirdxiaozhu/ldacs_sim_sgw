package system

import (
	"ldacs_sim_sgw/pkg/forward_module/f_config"
)

// 配置文件结构体
type System struct {
	Config f_config.ServerConfig `json:"config"`
}
