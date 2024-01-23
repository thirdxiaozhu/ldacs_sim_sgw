package handle

import (
	"encoding/binary"
	"errors"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"unsafe"

	"github.com/looplab/fsm"
)

// #cgo CFLAGS: -I /usr/local/include/ldacs
// #cgo LDFLAGS: -L /usr/local/lib/ldacs -lldacscore
// #include <ldacs_core/ldacs_core.h>
import "C"

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

	keyOcts := make([]uint8, 4)
	C.generate_rand((*C.uchar)(unsafe.Pointer(&keyOcts[0])))
	st.RandV = binary.BigEndian.Uint32(keyOcts)

	st.SharedKeyB = util.Base64Decode(st.SharedKey)
	st.KdfKB = make([]uint8, st.KdfLen)

	info := sharedInfo{
		Constant:     0x01,
		MacLen:       st.MacLen,
		AuthId:       st.AuthId,
		EncId:        st.EncId,
		RandV:        st.RandV,
		UaAs:         uint8(st.AsSac),
		UaGsc:        uint8(st.GscSac),
		KdfLen:       uint(st.KdfLen),
		SharedKeyLen: uint(len(st.SharedKeyB)),
	}

	e := C.generate_kdf_by_info((*C.struct_shared_info_s)(unsafe.Pointer(&info)), (*C.uchar)(unsafe.Pointer(&st.SharedKeyB[0])), (*C.uchar)(unsafe.Pointer(&st.KdfKB[0])))
	st.KdfK = util.Base64Encode(st.KdfKB)

	if e == 0 {
		return errors.New("fail")
	}

	return nil
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

type SecState struct {
	Name string
	FSM  *fsm.FSM
}

var (
	SecStates SecState
)
