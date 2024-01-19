package request

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		PREFIX_MODEL: global.PREFIX_MODEL{ID: 1},
		ParentId:     "0",
		Path:         "dashboard",
		Name:         "dashboard",
		Component:    "view/dashboard/index.vue",
		Sort:         1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
