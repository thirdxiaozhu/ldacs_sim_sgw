package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type AuthcStateSearch struct {
	StartCreatedAt      *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt        *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	AuthcAsSac          *int       `json:"authc_as_sac" form:"authc_as_sac" `
	AuthcGsSac          *int       `json:"authc_gs_sac" form:"authc_gs_sac" `
	AuthcGscSac         *int       `json:"authc_gsc_sac" form:"authc_gsc_sac" `
	AuthzState          *int       `json:"authz_state" form:"authz_state" `
	StartAuthcTransTime *time.Time `json:"startAuthcTransTime" form:"startAuthcTransTime"`
	EndAuthcTransTime   *time.Time `json:"endAuthcTransTime" form:"endAuthcTransTime"`
	request.PageInfo
}
