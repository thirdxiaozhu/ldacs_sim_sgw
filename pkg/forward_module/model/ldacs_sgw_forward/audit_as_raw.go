// 自动生成模板AuditAsRaw
package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"
)

// AS报文 结构体  AuditAsRaw
type AuditAsRaw struct {
	global.PREFIX_MODEL
	AuditAsSac int       `json:"audit_as_sac" form:"audit_as_sac" gorm:"column:audit_as_sac;comment:;"` //飞机站SAC
	AsSac      AccountAs `json:"as_sac" form:"as_sac" gorm:"foreignKey:AuditAsSac"`
	AuditAsMsg string    `json:"audit_as_msg" form:"audit_as_msg" gorm:"column:audit_as_msg;comment:;"` //飞机站报文
}

// TableName AS报文 AuditAsRaw自定义表名 audit_as_raw
func (AuditAsRaw) TableName() string {
	return "audit_as_raw"
}
