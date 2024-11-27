package global

type LdEnum interface {
	GetString() string
	CheckValid() bool
}

type AuthStateKind uint16

const (
	AUTH_STAGE_UNDEFINED AuthStateKind = iota /* 0 -- most likely accident */
	AUTH_STAGE_DELETING                       /* indicates state is dead but not yet freed */

	/* IKE states */
	AUTH_STAGE_A0
	AUTH_STAGE_G0
	AUTH_STAGE_A1
	AUTH_STAGE_G1
	AUTH_STAGE_A2
	AUTH_STAGE_G2

	AUTH_STAGE_OFF
)

func (f AuthStateKind) GetString() string {
	return [...]string{
		"AUTH_STAGE_UNDEFINED",
		"AUTH_STAGE_DELETING",
		"AUTH_STAGE_A0",
		"AUTH_STAGE_G0",
		"AUTH_STAGE_A1",
		"AUTH_STAGE_G1",
		"AUTH_STAGE_A2",
		"AUTH_STAGE_G2",
		"AUTH_STAGE_OFF",
	}[f]
}

func (f AuthStateKind) CheckValid() bool {
	return f <= AUTH_STAGE_OFF
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

type KeyLen uint8

const (
	AUTHC_KLEN_RESERVED KeyLen = 0x0
	AUTHC_KLEN_128      KeyLen = 0x1
	AUTHC_KLEN_256      KeyLen = 0x2
)

func (f KeyLen) GetString() string {
	return [...]string{
		"AUTHC_KLEN_RESERVED",
		"AUTHC_KLEN_128",
		"AUTHC_KLEN_256",
	}[f-AUTHC_KLEN_RESERVED]
}

func (f KeyLen) CheckValid() bool {
	return f <= AUTHC_KLEN_256
}

func (f KeyLen) GetKeyLen() uint {
	switch f {
	case AUTHC_KLEN_256:
		return 256
	case AUTHC_KLEN_128:
		return 128
	default:
		return 0
	}
}
