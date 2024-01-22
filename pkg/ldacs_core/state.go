package ldacscore

import "C"
import (
	"github.com/looplab/fsm"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
)

/* belong to AS_SAC */
type State struct {
	global.PREFIX_MODEL
	//AsSac     model.AccountAs `json:"as_sac"`
	SnpState  global.SnpStateKind  `json:"snp_state" gorm:"column:snp_state;type:int;default:0;"`
	AuthState global.AuthStateKind `json:"auth_state" gorm:"column:auth_state;type:int;default:0;"`
	IsTerm    bool                 `json:"is_term" gorm:"column:is_term;type:int;default:0;"`
	AsSac     uint8                `json:"as_sac" gorm:"column:as_sac;type:int;default:0;"`
	GsSac     uint8                `json:"gs_sac" gorm:"column:gs_sac;type:int;default:0;"`
	GscSac    uint8                `json:"gsc_sac" gorm:"column:gsc_sac;type:int;default:0;"`
	MacLen    uint8                `json:"mac_len" gorm:"column:mac_len;type:int;default:0;"`
	AuthId    uint8                `json:"auth_id" gorm:"column:auth_id;type:int;default:0;"`
	EncId     uint8                `json:"enc_id" gorm:"column:enc_id;type:int;default:0;"`
	RandV     uint32               `json:"rand_v" gorm:"column:rand_v;type:int;default:0;"`
	Sqn       uint32               `json:"sqn" gorm:"column:sqn;type:int;default:0;"`
	KdfLen    uint32               `json:"kdf_len" gorm:"column:kdf_len;type:int;default:0;"`
	SharedKey []uint8              `json:"shared_key" gorm:"column:shared_key;"`
	KdfK      []uint8              `json:"kdf_k" gorm:"column:kdf_k;"`
	IsSuccess uint8                `json:"is_success" gorm:"column:is_success;type:int;default:0;"`
	SecHead   SecHead              `json:"sec_head"`
	AuthFsm   fsm.FSM              `json:"auth_fsm"`
}

func initState(uas uint32) *State {
	st := State{
		SnpState:  global.SNP_STATE_CONNECTING,
		AuthState: global.AUTH_STATE_G0,
		IsTerm:    false,
		AsSac:     util.ParseUAs(uas, "AS"),
		GsSac:     util.ParseUAs(uas, "GS"),
		GscSac:    util.ParseUAs(uas, "GSC"),
		KdfLen:    19,
		SharedKey: util.GetShardKey(uas),
		AuthFsm:   *InitNewAuthFsm(),
	}

	return &st
}
