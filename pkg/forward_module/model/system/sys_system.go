package system

import (
	"ldacs_sim_sgw/pkg/forward_module/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
