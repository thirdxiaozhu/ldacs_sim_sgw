// 自动生成模板AccountAuthz
package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/pkg/forward_module/f_global"
)

// 业务权限 结构体  AccountAuthz
type AccountAuthz struct {
	f_global.GVA_MODEL
	Authz_name string `json:"authz_name" form:"authz_name" gorm:"column:authz_name;comment:;" binding:"required"` //权限名称
}

// TableName 业务权限 AccountAuthz自定义表名 account_authz
func (AccountAuthz) TableName() string {
	return "account_authz"
}
