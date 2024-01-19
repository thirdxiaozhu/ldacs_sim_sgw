package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type AuditAsRawSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	AuditAsSac     *int       `json:"audit_as_sac" form:"audit_as_sac" `
	AuditAsMsg     string     `json:"audit_as_msg" form:"audit_as_msg" `
	request.PageInfo
}
