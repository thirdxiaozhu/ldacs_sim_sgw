package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
)

type KeyEntitySearch struct {
	KeyID    string `json:"id" form:"id" `
	KeyType  string `json:"key_type" form:"key_type" gorm:"column:key_type;comment:;"`
	Owner1   string `json:"owner1" form:"owner1" gorm:"column:owner1;comment:AS_UA;"`
	Owner2   string `json:"owner2" form:"owner2" gorm:"column:owner2;comment:SGW_UA;"`
	KeyState string `json:"key_state" form:"key_state" gorm:"column:key_state;comment:;"`
	//Kind       string `json:"kind" form:"kind" `
	//User1      *int   `json:"user1" form:"user1" `
	//User2      *int   `json:"user2" form:"user2" `
	//KeyStatus  string `json:"key_status" form:"key_status"`
	//Ciphertext string `json:"ciphertext" form:"ciphertext" `
	request.PageInfo
}
