package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type AccountGsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	GsSac          *int       `json:"gs_sac" form:"gs_sac" `
	LatitudeN      *float64   `json:"latitude_n" form:"latitude_n" `
	LongtitudeE    *float64   `json:"longtitude_e" form:"longtitude_e" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
