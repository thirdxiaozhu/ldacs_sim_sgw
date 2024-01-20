// 自动生成模板AuthcState
package model

import (
	"ldacs_sim_sgw/internal/global"

	"time"
)

// 认证状态 结构体  AuthcState
type AuthcState struct {
	global.PREFIX_MODEL
	AuthcAsSac     int        `json:"authc_as_sac" form:"authc_as_sac" gorm:"column:authc_as_sac;comment:;"`    //认证飞机站
	AuthcGsSac     int        `json:"authc_gs_sac" form:"authc_gs_sac" gorm:"column:authc_gs_sac;comment:;"`    //当前地面站
	AuthcGscSac    int        `json:"authc_gsc_sac" form:"authc_gsc_sac" gorm:"column:authc_gsc_sac;comment:;"` //当前地面控制站
	AuthcState     int        `json:"authc_state" form:"authc_state" gorm:"column:authc_state;comment:;"`       //认证状态
	AsSac          AccountAs  `json:"as_sac" form:"as_sac" gorm:"foreignKey:AuthcAsSac;references:ID"`
	GsSac          AccountGs  `json:"gs_sac" form:"gs_sac" gorm:"foreignKey:AuthcGsSac;references:ID"`
	GscSac         AccountGsc `json:"gsc_sac" form:"gsc_sac" gorm:"foreignKey:AuthcGscSac;references:ID"`
	AuthcTransTime *time.Time `json:"authc_trans_time" form:"authc_trans_time" gorm:"column:authc_trans_time;comment:;"` //状态转换时间
}

// TableName 认证状态 AuthcState自定义表名 authc_state
func (AuthcState) TableName() string {
	return "authc_state"
}
