package util

import "ldacs_sim_sgw/internal/global"

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

func ParseUAs(uas uint32, tag string) uint8 {
	if tag == "AS" {
		return uint8(uas>>(global.UA_GS_LEN+global.UA_GSC_LEN)) & 0xFF
	} else if tag == "GS" {
		return uint8(uas>>global.UA_GSC_LEN) & 0x0F
	} else if tag == "GSC" {
		return uint8(uas) & 0x0F
	}
	return 0
}

func GenUAs(uaAs, uaGs, uaGsc uint8) uint32 {
	uaAs32 := uint32(uaAs)
	uaGs32 := uint32(uaGs)
	uaGsc32 := uint32(uaGsc)
	return uaAs32<<(global.UA_GS_LEN+global.UA_GSC_LEN) + uaGs32<<global.UA_GSC_LEN + uaGsc32
}

/*
未来应从数据库中读取
*/
func GetShardKey(uas uint32) []uint8 {
	keys := []uint8{0x12, 0x34, 0x56, 0x78}
	return keys[:]
}
