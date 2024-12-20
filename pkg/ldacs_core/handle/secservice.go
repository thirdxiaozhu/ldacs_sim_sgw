package handle

import "C"
import (
	"bytes"
	"crypto/rand"
	gmssl "github.com/thirdxiaozhu/GmSSL-Go"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/config"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
	"ldacs_sim_sgw/pkg/ldacs_core/service"
	"unsafe"
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
	ASSac  uint16        `ldacs:"name:as_sac; size:12; type:set"`
	GSSac  uint16        `ldacs:"name:gs_sac; size:12; type:set"`
	KeyLen global.KeyLen `ldacs:"name:key_len; size:2; type:enum"`
	N2     []byte        `ldacs:"name:N2; bytes_size: 16; type:fbytes"`
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
	KeyLen global.KeyLen `ldacs:"name:key_len; size:2; type:enum"`
	N2     []byte        `ldacs:"name:N2; bytes_size: 16; type:fbytes"`
}

var (
	distro string
)

const (
	KDF_ITER = 10000
)

func init() {
	distro = config.GetLinuxDistroCommand()
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

func CalcHMac(handler unsafe.Pointer, data []byte, limit uint32) (res []byte) {
	res, err := util.CalcHMAC(handler, data, limit)
	if err != nil {
		global.LOGGER.Error("Unable calc hmac", zap.Error(err))
		return nil
	}
	return
}

func VerifyHmac(handler unsafe.Pointer, data, toVerify []byte, limit uint32) bool {
	hmac := CalcHMac(handler, data, limit)

	return bytes.Equal(hmac, toVerify)
}

func GetKeyHandle(state, ktype string, owner1, owner2 uint32) (unsafe.Pointer, error) {
	dbName := global.CONFIG.Sqlite.Dsn()
	tableName := model.KeyEntity{}.TableName()

	km, err := service.KeyEntitySer.GetKeyEntityByContent(ldacs_sgw_forwardReq.KeyEntitySearch{
		KeyState: state,
		KeyType:  ktype,
		Owner1:   util.UAformat(owner1),
		Owner2:   util.UAformat(owner2),
	})

	handle, err := util.GetKeyHandle(dbName, tableName, km.KeyID)
	if err != nil {
		return nil, err
	}
	return handle, nil
}

func GenerateSharedKey(unit *LdacsUnit) (N2 []byte, err error) {
	st := unit.State
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

	unit.HandlerAsSgw, unit.KeyAsGs, err = SGWDeriveKey(util.UAformat(10010), util.UAformat(10086), util.UAformat(10000), uint32(SharedInfo.KeyLen.GetKeyLen()), random)

	if err != nil {
		return nil, err
	}
	return
}

func SGWDeriveKey(asUa, gsUa, sgwUa string, keyLen uint32, n []byte) (unsafe.Pointer, []byte, error) {

	dbName := global.CONFIG.Sqlite.Dsn()
	tableName := model.KeyEntity{}.TableName()

	km, err := service.KeyEntitySer.GetKeyEntityByContent(ldacs_sgw_forwardReq.KeyEntitySearch{
		KeyState: "ACTIVE",
		KeyType:  "ROOT_KEY",
		Owner1:   util.UAformat(asUa),
		Owner2:   util.UAformat(sgwUa),
	})

	if err != nil {
		return nil, nil, err
	}

	err = util.DeriveKey(dbName, tableName, km.KeyID, gsUa, keyLen, n)
	if err != nil {
		return nil, nil, err
	}

	mkeyAsSgw, err := service.KeyEntitySer.GetKeyEntityByContent(ldacs_sgw_forwardReq.KeyEntitySearch{
		KeyState: "ACTIVE",
		KeyType:  "MASTER_KEY_AS_SGW",
		Owner1:   util.UAformat(asUa),
		Owner2:   util.UAformat(sgwUa),
	})

	mkeyHandleAsSgw, err := util.GetKeyHandle(dbName, tableName, mkeyAsSgw.KeyID)
	if err != nil {
		return nil, nil, err
	}

	mkeyAsGs, err := service.KeyEntitySer.GetKeyEntityByContent(ldacs_sgw_forwardReq.KeyEntitySearch{
		KeyState: "ACTIVE",
		KeyType:  "MASTER_KEY_AS_GS",
		Owner1:   util.UAformat(asUa),
		Owner2:   util.UAformat(gsUa),
	})

	mkeyHandleAsGs, err := util.QueryKeyValue(dbName, tableName, mkeyAsGs.KeyID)
	if err != nil {
		return nil, nil, err
	}

	return mkeyHandleAsSgw, mkeyHandleAsGs.Key, nil
}
