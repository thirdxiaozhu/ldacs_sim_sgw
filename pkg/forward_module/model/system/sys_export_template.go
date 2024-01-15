// 自动生成模板SysExportTemplate
package system

import (
	"ldacs_sim_sgw/pkg/forward_module/forward_global"
)

// 导出模板 结构体  SysExportTemplate
type SysExportTemplate struct {
	forward_global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:模板名称;"`                       //模板名称
	TableName    string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:表名称;"`        //表名称
	TemplateID   string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:模板标识;"`    //模板标识
	TemplateInfo string `json:"templateInfo" form:"templateInfo" gorm:"column:template_info;type:text;"` //模板信息
}
