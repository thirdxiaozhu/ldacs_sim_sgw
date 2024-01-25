package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type AuthzPlaneSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Authz_planeId  *int       `json:"authz_PlaneId" form:"authz_PlaneId" `
	Authz_flight   *int       `json:"authz_flight" form:"authz_flight" `
	Authz_autz     *int       `json:"authz_autz" form:"authz_autz" `
	Authz_state    *int       `json:"authz_state" form:"authz_state" `
	request.PageInfo
}
