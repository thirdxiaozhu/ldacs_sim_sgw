// 自动生成模板KeyEntity
package model

import (
	"ldacs_sim_sgw/internal/global"
)

// 密钥 结构体  KeyEntity
type KeyEntity struct {
	global.PREFIX_MODEL
	KeyID      string `json:"key_id" form:"key_id" gorm:"column:key_id;comment:;"`                //ID
	Kind       string `json:"kind" form:"kind" gorm:"column:kind;comment:;"`                      //密钥类型
	User1      int    `json:"user1" form:"user1" gorm:"column:user1;comment:;"`                   //所有者1
	User2      int    `json:"user2" form:"user2" gorm:"column:user2;comment:;"`                   //所有者2
	Length     int    `json:"length" form:"length" gorm:"column:length;comment:;"`                //密钥长度
	KeyStatus  string `json:"key_status" form:"key_status" gorm:"column:key_status;comment:;"`    //密钥状态
	UpdateTime int    `json:"update_time" form:"update_time" gorm:"column:update_time;comment:;"` //更新间隔
	Ciphertext string `json:"ciphertext" form:"ciphertext" gorm:"column:ciphertext;comment:;"`    //密钥密文
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 密钥 KeyEntity自定义表名 key
func (KeyEntity) TableName() string {
	return "key"
}
