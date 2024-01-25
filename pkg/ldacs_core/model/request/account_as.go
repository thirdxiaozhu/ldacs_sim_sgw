package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type AccountAsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	AsPlaneId      *int       `json:"as_plane_id" form:"as_plane_id" `
	AsFlight       *int       `json:"as_flight" form:"as_flight" `
	StartAsDate    *time.Time `json:"startAsDate" form:"startAsDate"`
	EndAsDate      *time.Time `json:"endAsDate" form:"endAsDate"`
	AsSac          string     `json:"as_sac" form:"as_sac" `
	request.PageInfo
}
