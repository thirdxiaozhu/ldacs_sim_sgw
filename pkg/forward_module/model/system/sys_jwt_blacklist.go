package system

import (
	"ldacs_sim_sgw/pkg/forward_module/forward_global"
)

type JwtBlacklist struct {
	forward_global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
