package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
)

type KeyEntitySearch struct {
	KeyID      string `json:"id" form:"id" `
	Kind       string `json:"kind" form:"kind" `
	User1      *int   `json:"user1" form:"user1" `
	User2      *int   `json:"user2" form:"user2" `
	KeyStatus  string `json:"key_status" form:"key_status"`
	Ciphertext string `json:"ciphertext" form:"ciphertext" `
	request.PageInfo
}
