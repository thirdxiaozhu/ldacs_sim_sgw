package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type AccountPlaneSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Plane_id       string     `json:"plane_id" form:"plane_id" `
	Company        string     `json:"company" form:"company" `
	Model          string     `json:"model" form:"model" `
	request.PageInfo
}
