package handle

import "C"
import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	//"unsafe"
	//"ldacs_sim_sgw/internal/util"
	//"encoding/binary"
	//"errors"
)

//// #cgo CFLAGS: -I /usr/local/include/ldacs
//// #cgo LDFLAGS: -L /usr/local/lib/ldacs -lldacscore
//// #include <ldacs_core/ldacs_core.h>
//import "C"

type SEC_CMDS global.Constant

const (
	CMD_RESERVED SEC_CMDS = iota
	REGIONAL_BROADCASTING
	REGIONAL_ACCESS_REQ
	REGIONAL_ACCESS_RES
	REGIONAL_ACCESS_CONFIRM
	AS_STATUS_REQ
	AS_STATUS_RES
	SGW_STATUS_REQ
	SGW_STATUS_RES
	AUTH_FAILED
	STATUS_UPDATEED
	CMD_RESERVED_1
	CMD_RESERVED_2
	CMD_RESERVED_3
	CMD_RESERVED_4
	CMD_RESERVED_5
)

func (f SEC_CMDS) String() string {
	return [...]string{
		"CMD_RESERVED SEC_CMDS = iota",
		"REGIONAL_BROADCASTING",
		"REGIONAL_ACCESS_REQ",
		"REGIONAL_ACCESS_RES",
		"REGIONAL_ACCESS_CONFIRM",
		"AS_STATUS_REQ",
		"AS_STATUS_RES",
		"SGW_STATUS_REQ",
		"SGW_STATUS_RES",
		"AUTH_FAILED",
		"STATUS_UPDATEED",
		"CMD_RESERVED_1",
		"CMD_RESERVED_2",
		"CMD_RESERVED_3",
		"CMD_RESERVED_4",
		"CMD_RESERVED_5",
	}[f]
}

type sharedInfo struct {
	Constant     uint8
	MacLen       uint8
	AuthId       uint8
	EncId        uint8
	RandV        uint32
	UaAs         uint8
	UaGsc        uint8
	KdfLen       uint
	SharedKeyLen uint
}

func genSharedInfo(st *model.State) error {
	return nil
	//keyOcts := make([]uint8, 4)
	//C.generate_rand((*C.uchar)(unsafe.Pointer(&keyOcts[0])))
	//st.RandV = binary.BigEndian.Uint32(keyOcts)
	//
	//st.SharedKeyB = util.Base64Decode(st.SharedKey)
	//st.KdfKB = make([]uint8, st.KdfLen)
	//
	//info := sharedInfo{
	//	Constant:     0x01,
	//	MacLen:       st.MacLen,
	//	AuthId:       st.AuthId,
	//	EncId:        st.EncId,
	//	RandV:        st.RandV,
	//	UaAs:         uint8(st.AsSac),
	//	UaGsc:        uint8(st.GscSac),
	//	KdfLen:       uint(st.KdfLen),
	//	SharedKeyLen: uint(len(st.SharedKeyB)),
	//}
	//
	//e := C.generate_kdf_by_info((*C.struct_shared_info_s)(unsafe.Pointer(&info)), (*C.uchar)(unsafe.Pointer(&st.SharedKeyB[0])), (*C.uchar)(unsafe.Pointer(&st.KdfKB[0])))
	//st.KdfK = util.Base64Encode(st.KdfKB)
	//
	//if e == 0 {
	//	return errors.New("fail")
	//}
	//
	//return nil
}

type SecHead struct {
	Cmd     uint8 `json:"cmd"`
	Ver     uint8 `json:"ver"`
	ProId   uint8 `json:"pro_id"`
	Reserve uint8 `json:"reserve"`
}

type SecPldA1 struct {
	MacLen uint8 `json:"maclen"`
	AuthID uint8 `json:"authid"`
	EncID  uint8 `json:"encid"`
}
type SecPldKdf struct {
	MacLen uint8   `json:"maclen"`
	AuthID uint8   `json:"authid"`
	EncID  uint8   `json:"encid"`
	RandV  uint32  `json:"randv"`
	TimeV  uint64  `json:"time"`
	KdfKB  []uint8 `json:"kdfk"`
}
type SecPldKdfCon struct {
	IsOK int `json:"is_ok"`
}

/* ============================= */

type STYPE uint8

const (
	AUC_RQST     STYPE = 0x41
	AUC_RESP     STYPE = 0x42
	AUC_KEY_EXEC STYPE = 0x43
)

func (f STYPE) GetString() string {
	return [...]string{
		"AUC_RQST",
		"AUC_RESP",
		"AUC_KEY_EXEC",
	}[f-AUC_RQST]
}

func (f STYPE) CheckValid() bool {
	return f >= AUC_RQST && f <= AUC_KEY_EXEC
}

type PID uint8

const (
	PID_RESERVED PID = 0x0
	PID_SIGN     PID = 0x1
	PID_MAC      PID = 0x2
	PID_BOTH     PID = 0x3
)

func (f PID) GetString() string {
	return [...]string{
		"PID_RESERVED",
		"PID_SIGN",
		"PID_MAC",
		"PID_BOTH",
	}[f-PID_RESERVED]
}

func (f PID) CheckValid() bool {
	return f <= PID_BOTH
}

type MacLen uint8

const (
	AUTHC_MACLEN_INVALID MacLen = 0x0
	AUTHC_MACLEN_96      MacLen = 0x1
	AUTHC_MACLEN_128     MacLen = 0x2
	AUTHC_MACLEN_64      MacLen = 0x3
	AUTHC_MACLEN_256     MacLen = 0x4
)

func (f MacLen) GetString() string {
	return [...]string{
		"AUTHC_MACLEN_INVALID",
		"AUTHC_MACLEN_96",
		"AUTHC_MACLEN_128",
		"AUTHC_MACLEN_64",
		"AUTHC_MACLEN_256",
	}[f-AUTHC_MACLEN_INVALID]
}

func (f MacLen) CheckValid() bool {
	return f <= AUTHC_MACLEN_256
}

type AuthID uint8

const (
	AUTHC_AUTH_INVALID      AuthID = 0x0
	AUTHC_AUTH_SM3HMAC      AuthID = 0x1
	AUTHC_AUTH_SM2_WITH_SM3 AuthID = 0x2
)

func (f AuthID) GetString() string {
	return [...]string{
		"AUTHC_AUTH_INVALID",
		"AUTHC_AUTH_SM3HMAC",
		"AUTHC_AUTH_SM2_WITH_SM3",
	}[f-AUTHC_AUTH_INVALID]
}

func (f AuthID) CheckValid() bool {
	return f <= AUTHC_AUTH_SM2_WITH_SM3
}

type EncID uint8

const (
	AUTHC_ENC_INVALID EncID = 0x0
	AUTHC_ENC_SM4_CBC EncID = 0x1
	AUTHC_ENC_SM4_CFB EncID = 0x2
	AUTHC_ENC_SM4_OFB EncID = 0x3
	AUTHC_ENC_SM4_ECB EncID = 0x4
	AUTHC_ENC_SM4_CTR EncID = 0x5
)

func (f EncID) GetString() string {
	return [...]string{
		"AUTHC_ENC_INVALID",
		"AUTHC_ENC_SM4_CBC",
		"AUTHC_ENC_SM4_CFB",
		"AUTHC_ENC_SM4_OFB",
		"AUTHC_ENC_SM4_ECB",
		"AUTHC_ENC_SM4_CTR",
	}[f-AUTHC_ENC_INVALID]
}

func (f EncID) CheckValid() bool {
	return f <= AUTHC_ENC_SM4_CTR
}

type AucRqst struct {
	Stype  STYPE  `ldacs:"name:S_TYPE; size:8; type:enum"`
	Ver    uint8  `ldacs:"name:VER; size:3; type:set"`
	PID    PID    `ldacs:"name:PID; size:2; type:enum"`
	ASSac  uint16 `ldacs:"name:as_sac; size:12; type:set"`
	GSSac  uint16 `ldacs:"name:gs_sac; size:12; type:set"`
	MacLen MacLen `ldacs:"name:maclen; size:4; type:enum"`
	AuthID AuthID `ldacs:"name:authid; size:4; type:enum"`
	EncID  EncID  `ldacs:"name:encid; size:4; type:enum"`
	N1     []byte `ldacs:"name:N1; bytes_size: 16; type:fbytes"`
}
