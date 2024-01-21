package ldacscore

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"
	"unsafe"

	"github.com/looplab/fsm"
)

// #cgo CFLAGS: -I /usr/local/include/ldacs
// #cgo LDFLAGS: -L /usr/local/lib/ldacs -lldacscore
// #include <ldacs_core/ldacs_core.h>
import "C"

type SEC_CMDS Constant

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

func genSharedInfo(st *state) error {

	keyOcts := make([]uint8, 4)
	C.generate_rand((*C.uchar)(unsafe.Pointer(&keyOcts[0])))
	st.RandV = binary.BigEndian.Uint32(keyOcts)

	info := sharedInfo{
		Constant:     0x01,
		MacLen:       st.MacLen,
		AuthId:       st.AuthId,
		EncId:        st.EncId,
		RandV:        st.RandV,
		UaAs:         st.UaAs,
		UaGsc:        st.UaGsc,
		KdfLen:       uint(st.KdfLen),
		SharedKeyLen: uint(len(st.SharedKey)),
	}

	st.KdfK = make([]uint8, st.KdfLen)
	e := C.generate_kdf_by_info((*C.struct_shared_info_s)(unsafe.Pointer(&info)), (*C.uchar)(unsafe.Pointer(&st.SharedKey[0])), (*C.uchar)(unsafe.Pointer(&st.KdfK[0])))

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
	KdfK   []uint8 `json:"kdfk"`
}
type SecPldKdfCon struct {
	IsOK uint8 `json:"is_ok"`
}

type SecState struct {
	Name string
	FSM  *fsm.FSM
}

var (
	SecStates SecState
)

func (s *SecState) beforeAuthStateG0(ctx context.Context, e *fsm.Event) error {
	return nil
}
func (s *SecState) beforeAuthStateG1(ctx context.Context, e *fsm.Event) error {
	node := ctx.Value("node").(*LdacsStateConnNode)
	st := node.State

	st.SecHead.Cmd = uint8(REGIONAL_ACCESS_RES)

	err := genSharedInfo(st)
	if err != nil {
		return err
	}

	kdfMsg := SecPldKdf{
		MacLen: st.MacLen,
		AuthID: st.AuthId,
		EncID:  st.EncId,
		RandV:  st.RandV,
		TimeV:  uint64(time.Now().Unix()),
		KdfK:   st.KdfK,
	}
	unitData, _ := json.Marshal(kdfMsg)

	pktUnit := LdacsUnit{
		AsSac: st.UaAs,
		UaGs:  st.UaGs,
		UaGsc: st.UaGsc,
		Head:  st.SecHead,
		Data:  unitData,
	}

	node.ToSendPkt(&pktUnit)

	return nil
}
func (s *SecState) beforeAuthStateG2(ctx context.Context, e *fsm.Event) error {
	return nil
}

func (s *SecState) beforeAuthStateUndef(ctx context.Context, e *fsm.Event) error {

	return nil
}

func getFSMEvents(dst string, src ...string) *fsm.EventDesc {
	return &fsm.EventDesc{
		Name: dst,
		Src:  src,
		Dst:  dst,
	}
}

func (s *SecState) handleErrEvent(ctx context.Context, err error) {
	if err != nil {
		err := s.FSM.Event(ctx, AUTH_STATE_UNDEFINED.String())
		if err != nil {
			return
		}
	}
}

func InitNewAuthFsm() *fsm.FSM {
	return fsm.NewFSM(AUTH_STATE_UNDEFINED.String(),
		fsm.Events{
			*getFSMEvents(AUTH_STATE_G0.String(), AUTH_STATE_UNDEFINED.String()),
			*getFSMEvents(AUTH_STATE_G1.String(), AUTH_STATE_G0.String()),
			*getFSMEvents(AUTH_STATE_G2.String(), AUTH_STATE_G1.String()),

			//处理错误
			*getFSMEvents(AUTH_STATE_UNDEFINED.String(), AUTH_STATE_G0.String(), AUTH_STATE_G1.String(), AUTH_STATE_G2.String(), AUTH_STATE_UNDEFINED.String()),
		},
		fsm.Callbacks{
			"before_" + AUTH_STATE_G0.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG0(ctx, e))
			},
			"before_" + AUTH_STATE_G1.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG1(ctx, e))
			},
			"before_" + AUTH_STATE_G2.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG2(ctx, e))
			},
			"before_" + AUTH_STATE_UNDEFINED.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateUndef(ctx, e))
			},
		},
	)
}
