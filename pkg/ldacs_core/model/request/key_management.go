package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"time"
)

type KeyEntitySearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	KeyID          string     `json:"key_id" form:"key_id" `
	Kind           string     `json:"kind" form:"kind" `
	User1          *int       `json:"user1" form:"user1" `
	User2          *int       `json:"user2" form:"user2" `
	KeyStatus      string     `json:"key_status" form:"key_status"`
	Ciphertext     string     `json:"ciphertext" form:"ciphertext" `
	request.PageInfo
}
