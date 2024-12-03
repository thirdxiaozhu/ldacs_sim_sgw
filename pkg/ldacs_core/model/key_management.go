// 自动生成模板KeyEntity
package model

// 密钥 结构体  KeyEntity
type KeyEntity struct {
	//global.PREFIX_MODEL
	KeyID       string `json:"id" form:"id" gorm:"column:id;comment:;"`
	KeyType     string `json:"key_type" form:"key_type" gorm:"column:key_type;comment:;"`
	Owner1      string `json:"owner1" form:"owner1" gorm:"column:owner1;comment:AS_UA;"`
	Owner2      string `json:"owner2" form:"owner2" gorm:"column:owner2;comment:SGW_UA;"`
	KeyCipher   string `json:"key_cipher" form:"key_cipher" gorm:"column:key_cipher;comment:;"`
	KeyLen      int    `json:"key_len" form:"key_len" gorm:"column:key_len;comment:;"`
	KeyState    string `json:"key_state" form:"key_state" gorm:"column:key_state;comment:;"`
	CreateTime  string `json:"creatime" form:"creatime" gorm:"column:creatime;comment:;"`
	UpdateCycle int    `json:"updatecycle" form:"updatecycle" gorm:"column:updatecycle;comment:;"`
	KekLen      int    `json:"kek_len" form:"kek_len" gorm:"column:kek_len;comment:;"`
	KekCipher   string `json:"kek_cipher" form:"kek_cipher" gorm:"column:kek_cipher;comment:;"`
	Iv          string `json:"iv" form:"iv" gorm:"column:iv;comment:;"`
	IvLen       int    `json:"iv_len" form:"iv_len" gorm:"column:iv_len;comment:;"`
	ChckAlgo    string `json:"chck_algo" form:"chck_algo" gorm:"column:chck_algo;comment:;"`
	ChckLen     int    `json:"check_len" form:"check_len" gorm:"column:check_len;comment:;"`
	ChckValue   string `json:"chck_value" form:"chck_value" gorm:"column:chck_value;comment:;"`
	UpdateCount int    `json:"update_count" form:"update_count" gorm:"column:update_count;comment:;"`
}

// TableName 密钥 KeyEntity自定义表名 key
func (KeyEntity) TableName() string {
	return "sgw_keystore"
}
