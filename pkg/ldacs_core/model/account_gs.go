// 自动生成模板AccountGs
package model

import (
	"ldacs_sim_sgw/internal/global"
)

// 地面站 结构体  AccountGs
type AccountGs struct {
	global.PREFIX_MODEL
	GsSac       int     `json:"gs_sac" form:"gs_sac" gorm:"column:gs_sac;comment:;"`                   //地面站SAC
	LatitudeN   float64 `json:"latitude_n" form:"latitude_n" gorm:"column:latitude_n;comment:;"`       //北纬
	LongtitudeE float64 `json:"longtitude_e" form:"longtitude_e" gorm:"column:longtitude_e;comment:;"` //东经
}

// TableName 地面站 AccountGs自定义表名 account-gs
func (AccountGs) TableName() string {
	return "account-gs"
}
