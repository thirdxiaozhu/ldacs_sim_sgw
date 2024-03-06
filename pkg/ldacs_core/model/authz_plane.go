// 自动生成模板AuthzPlane
package model

import (
	"ldacs_sim_sgw/internal/global"
)

type AuthzPlaneMulti struct {
	AuthzAs     int   `json:"authz_as" `
	AuthzAuthzs []int `json:"authz_authz" `
}

// 飞机业务授权 结构体  AuthzPlane
type AuthzPlane struct {
	global.PREFIX_MODEL
	AuthzAs    int          `json:"authz_as" form:"authz_as" gorm:"column:authz_as;comment:;"`          //被授权飞机
	AuthzAuthz int          `json:"authz_autz" form:"authz_autz" gorm:"column:authz_autz;comment:;"`    //权限
	AuthzState int          `json:"authz_state" form:"authz_state" gorm:"column:authz_state;comment:;"` //授权状态
	AccountAs  AccountAs    `json:"as" form:"as" gorm:"foreignKey:AuthzAs;references:ID;"`
	Authz      AccountAuthz `json:"authz" form:"authz" gorm:"foreignKey:AuthzAuthz;references:ID;"`
}

type AuthzOptions struct {
	AuthzPlaneIds []AccountPlane  `json:"plane_ids"`
	AuthzFlights  []AccountFlight `json:"flights"`
	AuthzAuthzs   []AccountAuthz  `json:"authzs"`
}

// TableName 飞机业务授权 AuthzPlane自定义表名 authz_plane
func (AuthzPlane) TableName() string {
	return "authz_plane"
}
