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
	KUPDATE_REQUEST STYPE = 0x44
	KUPDATE_RESPONSE STYPE = 0x45
)

func (f STYPE) GetString() string {
	return [...]string{
		"AUC_RQST",
		"AUC_RESP",
		"AUC_KEY_EXEC",
		"KUPDATE_REQUEST",
		"KUPDATE_RESPONSE",
	}[f-STYPE-INVALID] // check：modified
}

func (f STYPE) CheckValid() bool {
	return f >= AUC_RQST && f <= KUPDATE_RESPONSE 
}

/*----------------------------------------------------------------*/
type GTYPE uint8

const (
	KUPDATE_REMIND GTYPE = 0x03 // check：检查编码
	KUPDATE_REQUEST GTYPE = 0x04
	KUPDATE_RESPONSE GTYPE = 0x05
	KEY_TRANSPORT GTYPE = 0x06 
)

func (f G_TYPE) GetString() string {
	return [...]string{
		"KUPDATE_REMIND", 
		"KUPDATE_REQUEST".
		"KUPDATE_RESPONSE",
		"KEY_TRANSPORT", 
	}[f-GTYPE-INVALID]
}

func (f G_TYPE) CheckValid() bool {
	return f >= KEY_TRANSPORT && f <= KUPDATE_REMIND 
}
/*----------------------------------------------------------------*/

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
/*----------------------------------------------------------------*/

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

func (f MacLen) GetMacLen() uint32 {
	switch f {
	case AUTHC_MACLEN_64:
		return 64 >> 3
	case AUTHC_MACLEN_96:
		return 96 >> 3
	case AUTHC_MACLEN_128:
		return 128 >> 3
	case AUTHC_MACLEN_256:
		return 256 >> 3
	default:
		return 0
	}
}
/*----------------------------------------------------------------*/

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
/*----------------------------------------------------------------*/

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
/*----------------------------------------------------------------*/

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
		return 256 >> 3
	case AUTHC_KLEN_128:
		return 128 >> 3
	default:
		return 0
	}
}
/*----------------------------------------------------------------*/

type ElementType uint8

const (
	KEY_TRANPORT_AFTER_AKA ElementType = 0x08
	KEY_UPDATE_KTRANPORT ElementType = 0x09
	KEY_UPDATE_REMIND ElementType = 0x0c
)

func (f ElementType) GetString() string {
	return [...]string{
		KEY_TRANPORT_AFTER_AKA: "KEY_TRANPORT_AFTER_AKA",
		KEY_UPDATE_KTRANPORT:   "KEY_UPDATE_KTRANPORT",
		KEY_UPDATE_REMIND:      "KEY_UPDATE_REMIND",
	}[f-ELEMENTTYPE-INVALID]
}

func (f ElementType) CheckValid() bool {
	return f >= KEY_TRANPORT_AFTER_AKA && f <= KEY_UPDATE_REMIND
}

/*----------------------------------------------------------------*/

type KeyType uint8

const (
	MASTER_KEY_AS_GS_128 KeyType = 0x01
	MASTER_KEY_AS_SGW_128 KeyType = 0x02
	MASTER_KEY_AS_SGW_256 KeyType = 0x03
)

func (k KeyType)GetString() string {
    return [...]string{
        MASTER_KEY_AS_GS_128:    "MASTER_KEY_AS_GS_128",
        MASTER_KEY_AS_SGW_128:   "MASTER_KEY_AS_SGW_128",
        MASTER_KEY_AS_SGW_256:   "MASTER_KEY_AS_SGW_256",
    }[f-KEYTYPE-INVALID]
}
func (k KeyType) CheckValid() bool {
    return k >= MASTER_KEY_AS_GS_128 && k <= MASTER_KEY_AS_SGW_256
}
/*----------------------------------------------------------------*/
