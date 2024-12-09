package model

import "C"
import (
	"ldacs_sim_sgw/internal/global"
)

/* belong to AS_SAC */
type State struct {
	global.PREFIX_MODEL
	AsSac      uint16               `json:"as_sac" form:"as_sac" gorm:"column:as_sac;comment:;"`
	GsSac      uint16               `json:"gs_sac" form:"gs_sac" gorm:"column:gs_sac;comment:;"`
	GscSac     uint16               `json:"gsc_sac" form:"gsc_sac" gorm:"column:gsc_sac;comment:;"`
	AccountGs  AccountGs            `json:"account_gs" form:"account_gs" gorm:"foreignKey:GsSac;references:GsSac"`
	AccountGsc AccountGsc           `json:"account_gsc" form:"account_gsc" gorm:"foreignKey:GscSac;references:GscSac"`
	AuthState  global.AuthStateKind `json:"auth_state" form:"auth_state" gorm:"column:auth_state;type:int;default:0;"`
	IsTerm     int                  `json:"is_term" form:"is_term" gorm:"column:is_term;type:int;default:0;"`
	MacLen     uint8                `json:"mac_len" form:"mac_len" gorm:"column:mac_len;type:int;default:0;"`
	AuthId     uint8                `json:"auth_id" form:"auth_id"  gorm:"column:auth_id;type:int;default:0;"`
	EncId      uint8                `json:"enc_id" form:"enc_id" gorm:"column:enc_id;type:int;default:0;"`
	Ver        uint8                `json:"version" form:"version" gorm:"column:version;type:int;default:0;"`
	PID        uint8                `json:"pid" form:"pid" gorm:"column:pid;type:int;default:0;"`
	RandV      uint32               `json:"rand_v" form:"rand_v" gorm:"column:rand_v;type:uint;default:0;"`
	Sqn        uint32               `json:"sqn"  form:"sqn" gorm:"column:sqn;type:int;default:0;"`
	KdfLen     uint8                `json:"kdf_len" form:"kdf_len"  gorm:"column:kdf_len;type:int;default:0;"`
	SharedKey  string               `json:"shared_key" form:"shared_key" gorm:"column:shared_key;"`
	KdfKB      []uint8              `gorm:"-"`
	IsSuccess  int                  `json:"is_success" form:"is_success" gorm:"column:is_success;type:int;default:0;"`
}

func NewState() *State {
	return &State{
		AuthState: global.AUTH_STAGE_UNDEFINED,
		KdfLen:    uint8(global.AUTHC_KLEN_128),
		IsTerm:    0,
	}

}

func (State) TableName() string {
	return "state"
}

func (State) TableNameU() string {
	return "State"
}
