package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
