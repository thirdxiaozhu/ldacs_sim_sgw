package test

import (
	"github.com/hdt3213/godis/lib/logger"
	"ldacs_sim_sgw/internal/util"
	"testing"
)

type STYPE uint8

const (
	AUC_RQST     STYPE = 0x41
	AUC_RESP     STYPE = 0x42
	AUC_KEY_EXEC STYPE = 0x43
	KUPDATE_REMIND GTYPE = 0x44 // check : how to handle gtype
	KUPDATE_REQUEST STYPE = 0x45
	KUPDATE_RESPONSE STYPE = 0x46
	KEY_TRANSPORT GTYPE = 0x47 // check : how to handle gtype
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

type AucRqst struct {
	Stype STYPE  `ldacs:"name:S_TYPE; size:8; type:enum"`
	Ver   uint8  `ldacs:"name:VER; size:3; type:set"`
	PID   PID    `ldacs:"name:PID; size:2; type:enum"`
	N1    []byte `ldacs:"name:N1; bytes_size: 3; type:fbytes"`
	//N1 []byte `ldacs:"name:N1; type:dbytes"`
	//ASSac  uint16 `ldacs:"required;minlength:8"`
	//GSSac  uint16 `ldacs:"required;minlength:8"`
	//MacLen uint8  `ldacs:"required;minlength:8"`
	//AuthID uint8  `ldacs:"required;minlength:8"`
	//EncID  uint8  `ldacs:"required;minlength:8"`
}

func TestTag(t *testing.T) {
	rqst := AucRqst{
		Stype: 0x41,
		Ver:   0x1,
		PID:   0x1,
		N1:    []byte{0x01, 0x02, 0x03},
	}
	bytes, err := util.MarshalLdacsPkt(rqst)

	if err != nil {
		logger.Error("Validation failed:", err)
	} else {
		logger.Info("Validation succeeded")
	}

	logger.Warn("Marshaled:", bytes)

	rqst2 := AucRqst{
		N1: make([]byte, 3),
	}
	_, err = util.UnmarshalLdacsPkt(bytes, &rqst2)
	if err != nil {
		logger.Error(err)
	}

	logger.Warn("Unmarshaled: ", rqst2)
}
