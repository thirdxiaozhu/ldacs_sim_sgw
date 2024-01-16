package system

import (
	"ldacs_sim_sgw/pkg/forward_module/f_global"
)

type JwtBlacklist struct {
	f_global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
