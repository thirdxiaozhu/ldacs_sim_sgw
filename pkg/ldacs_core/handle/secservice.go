package handle

import "C"
import (
	"crypto/rand"
	gmssl "github.com/GmSSL/GmSSL-Go"
	"github.com/hdt3213/godis/lib/logger"
	"ldacs_sim_sgw/internal/config"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
)

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

type AucRqst struct {
	Stype  global.STYPE  `ldacs:"name:S_TYPE; size:8; type:enum"`
	Ver    uint8         `ldacs:"name:VER; size:3; type:set"`
	PID    global.PID    `ldacs:"name:PID; size:2; type:enum"`
	ASSac  uint16        `ldacs:"name:as_sac; size:12; type:set"`
	GSSac  uint16        `ldacs:"name:gs_sac; size:12; type:set"`
	MacLen global.MacLen `ldacs:"name:maclen; size:4; type:enum"`
	AuthID global.AuthID `ldacs:"name:authid; size:4; type:enum"`
	EncID  global.EncID  `ldacs:"name:encid; size:4; type:enum"`
	N1     []byte        `ldacs:"name:N1; bytes_size: 16; type:fbytes"`
}

type AucSharedInfo struct {
	MacLen global.MacLen `ldacs:"name:maclen; size:4; type:enum"`
	AuthID global.AuthID `ldacs:"name:authid; size:4; type:enum"`
	EncID  global.EncID  `ldacs:"name:encid; size:4; type:enum"`
	N2     []byte        `ldacs:"name:N2; bytes_size: 16; type:fbytes"`
	ASSac  uint16        `ldacs:"name:as_sac; size:12; type:set"`
	GSSac  uint16        `ldacs:"name:gs_sac; size:12; type:set"`
	KeyLen global.KeyLen `ldacs:"name:key_len; size:2; type:enum"`
}

type AucResp struct {
	Stype  global.STYPE  `ldacs:"name:S_TYPE; size:8; type:enum"`
	Ver    uint8         `ldacs:"name:VER; size:3; type:set"`
	PID    global.PID    `ldacs:"name:PID; size:2; type:enum"`
	ASSac  uint16        `ldacs:"name:as_sac; size:12; type:set"`
	GSSac  uint16        `ldacs:"name:gs_sac; size:12; type:set"`
	MacLen global.MacLen `ldacs:"name:maclen; size:4; type:enum"`
	AuthID global.AuthID `ldacs:"name:authid; size:4; type:enum"`
	EncID  global.EncID  `ldacs:"name:encid; size:4; type:enum"`
	N2     []byte        `ldacs:"name:N2; bytes_size: 16; type:fbytes"`
	KeyLen global.KeyLen `ldacs:"name:key_len; size:2; type:enum"`
}

var (
	distro string
)

const (
	KDF_ITER = 1024
)

func init() {
	distro = config.GetLinuxDistroCommand()
	logger.Warn(distro)
}

func GenerateRandomBytes(size uint) []byte {
	switch distro {
	case "Ubuntu":
		key, err := gmssl.RandBytes(16)
		if err != nil {
			return nil
		}
		return key
	case "CentOS":

		return nil
	default:
		key := make([]byte, size)
		_, err := rand.Read(key)
		if err != nil {
			return nil
		}
		return key
	}
}

func GenerateSharedKey(st *model.State) (key, N2 []byte, err error) {
	N2 = GenerateRandomBytes(16)

	SharedInfo := AucSharedInfo{
		MacLen: global.MacLen(st.MacLen),
		AuthID: global.AuthID(st.AuthId),
		EncID:  global.EncID(st.EncId),
		N2:     N2,
		ASSac:  uint16(st.AsSac),
		GSSac:  uint16(st.GsSac),
		KeyLen: global.KeyLen(st.KdfLen),
	}

	random, err := util.MarshalLdacsPkt(SharedInfo)
	if err != nil {
		return nil, nil, err
	}

	logger.Warn(string(random))
	var salt [32]byte
	pbkdf2, err := gmssl.Sm3Pbkdf2(string(random), salt[:], KDF_ITER, uint(SharedInfo.KeyLen))
	if err != nil {
		return nil, nil, err
	}
	return pbkdf2, N2, nil
}
