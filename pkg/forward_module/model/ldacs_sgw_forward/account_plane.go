package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"
)

// AccountPlane 飞机账户管理 结构体  AccountPlane
type AccountPlane struct {
	global.PREFIX_MODEL
	PlaneId string `json:"plane_id" form:"plane_id" gorm:"column:plane_id;comment:;" binding:"required"` //飞机注册号
	Company string `json:"company" form:"company" gorm:"column:company;comment:;"`                       //所属航司
	Model   string `json:"model" form:"model" gorm:"column:model;comment:;" binding:"required"`          //飞机型号
	UA      string `json:"ua" form:"ua" gorm:"column:ua;default:0;not null;unique;"`
	TestUA  int    `json:"test_ua" form:"test_ua" gorm:"column:test_ua;unique;"`
}

// TableName 飞机账户管理 AccountPlane自定义表名 account_plane
func (AccountPlane) TableName() string {
	return "account_plane"
}
