package system

import (
	"ldacs_sim_sgw/internal/global"
)

type JwtBlacklist struct {
	global.PREFIX_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
