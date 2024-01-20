// 自动生成模板AccountGsc
package model

import (
	"ldacs_sim_sgw/internal/global"
)

// 地面控制站 结构体  AccountGsc
type AccountGsc struct {
	global.PREFIX_MODEL
	GscSac *int `json:"gsc_sac" form:"gsc_sac" gorm:"column:gsc_sac;comment:;"binding:"required"` //地面控制站SAC
}

// TableName 地面控制站 AccountGsc自定义表名 account-gsc
func (AccountGsc) TableName() string {
	return "account-gsc"
}
