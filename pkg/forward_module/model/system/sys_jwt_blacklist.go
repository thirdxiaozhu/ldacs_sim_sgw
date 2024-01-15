package system

import (
	"ldacs_sim_sgw/pkg/forward_module/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
