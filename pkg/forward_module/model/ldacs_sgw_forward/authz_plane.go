// 自动生成模板AuthzPlane
package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/pkg/forward_module/f_global"
)

type AuthzPlaneMulti struct {
	AuthzPlaneId int   `json:"authz_PlaneId"  `
	AuthzFlight  int   `json:"authz_flight" `
	AuthzAuthzs  []int `json:"authz_authz" `
}

// 飞机业务授权 结构体  AuthzPlane
type AuthzPlane struct {
	f_global.GVA_MODEL
	AuthzPlaneId int           `json:"authz_PlaneId" form:"authz_PlaneId" gorm:"column:authz_plane_id;comment:;"` //被授权飞机
	AuthzFlight  int           `json:"authz_flight" form:"authz_flight" gorm:"column:authz_flight;comment:;"`     //被授权航班
	AuthzAuthz   int           `json:"authz_autz" form:"authz_autz" gorm:"column:authz_autz;comment:;"`           //权限
	AuthzState   int           `json:"authz_state" form:"authz_state" gorm:"column:authz_state;comment:;"`        //授权状态
	Planeid      AccountPlane  `json:"plane_id" form:"plane_id" gorm:"foreignKey:AuthzPlaneId"`
	Flight       AccountFlight `json:"flight" form:"flight" gorm:"foreignKey:AuthzFlight"`
	Authz        AccountAuthz  `json:"authz" form:"authz" gorm:"foreignKey:AuthzAuthz"`
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
