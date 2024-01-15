package system

import "ldacs_sim_sgw/pkg/forward_module/forward_global"

type SysBaseMenuBtn struct {
	forward_global.GVA_MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
