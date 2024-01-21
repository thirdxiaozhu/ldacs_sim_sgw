package ldacscore

import "C"
import (
	"github.com/looplab/fsm"
)

type state struct {
	SnpState  snpStateKind
	AuthState authStateKind
	IsTerm    bool
	UaAs      uint8
	UaGs      uint8
	UaGsc     uint8
	MacLen    uint8
	AuthId    uint8
	EncId     uint8
	RandV     uint32
	Sqn       uint32
	KdfLen    uint32
	SharedKey []uint8
	KdfK      []uint8
	IsOK      uint8
	SecHead   SecHead
	AuthFsm   fsm.FSM
}

//func (s *sharedInfo) Pack(out unsafe.Pointer) {
//	buf := &bytes.Buffer{}
//	binary.Write(buf, binary.LittleEndian, s.Constant)
//	binary.Write(buf, binary.LittleEndian, s.MacLen)
//	binary.Write(buf, binary.LittleEndian, s.AuthId)
//	binary.Write(buf, binary.LittleEndian, s.EncId)
//	binary.Write(buf, binary.LittleEndian, s.RandV)
//	binary.Write(buf, binary.LittleEndian, s.UaAs)
//	binary.Write(buf, binary.LittleEndian, s.UaGsc)
//	binary.Write(buf, binary.LittleEndian, s.KdfLen)
//	binary.Write(buf, binary.LittleEndian, s.SharedKeyLen)
//	//Getting the lenfth of memory
//	l := buf.Len()
//	//Cast the point to byte slie to allow for direct memory manipulatios
//	o := (*[1 << 20]C.uchar)(out)
//	//Write to memory
//	for i := 0; i < l; i++ {
//		b, _ := buf.ReadByte()
//		o[i] = C.uchar(b)
//	}
//}

func parseUAs(uas uint32, tag string) uint8 {
	if tag == "AS" {
		return uint8(uas>>(UA_GS_LEN+UA_GSC_LEN)) & 0xFF
	} else if tag == "GS" {
		return uint8(uas>>UA_GSC_LEN) & 0x0F
	} else if tag == "GSC" {
		return uint8(uas) & 0x0F
	}
	return 0
}

func genUAs(uaAs, uaGs, uaGsc uint8) uint32 {
	uaAs32 := uint32(uaAs)
	uaGs32 := uint32(uaGs)
	uaGsc32 := uint32(uaGsc)
	return uaAs32<<(UA_GS_LEN+UA_GSC_LEN) + uaGs32<<UA_GSC_LEN + uaGsc32
}

/*
未来应从数据库中读取
*/
func GetShardKey(uas uint32) []uint8 {
	keys := []uint8{0x12, 0x34, 0x56, 0x78}
	return keys[:]
}

func initState(uas uint32) *state {
	st := state{
		SnpState:  SNP_STATE_CONNECTING,
		AuthState: AUTH_STATE_G0,
		IsTerm:    false,
		UaAs:      parseUAs(uas, "AS"),
		UaGs:      parseUAs(uas, "GS"),
		UaGsc:     parseUAs(uas, "GSC"),
		KdfLen:    19,
		SharedKey: GetShardKey(uas),
		AuthFsm:   *InitNewAuthFsm(),
	}

	return &st
}
