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

func ParseUAs(uas uint64, tag string) uint64 {
	if tag == "AS" {
		return uint64(uas>>(global.UA_GS_LEN+global.UA_GSC_LEN)) & 0xFF
	} else if tag == "GS" {
		return uint64(uas>>global.UA_GSC_LEN) & 0x0F
	} else if tag == "GSC" {
		return uint64(uas) & 0x0F
	}
	return 0
}

func GenUAs(uaAs, uaGs, uaGsc uint64) uint64 {
	uaAs32 := uint64(uaAs)
	uaGs32 := uint64(uaGs)
	uaGsc32 := uint64(uaGsc)
	return uaAs32<<(global.UA_GS_LEN+global.UA_GSC_LEN) + uaGs32<<global.UA_GSC_LEN + uaGsc32
}

/*
未来应从数据库中读取
*/
func GetShardKey(as_sac uint64) string {
	keys := []uint8{0x12, 0x34, 0x56, 0x78}
	return Base64Encode(keys[:])
}
