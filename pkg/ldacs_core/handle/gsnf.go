package handle

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
)

type GsnfPkt struct {
	GType uint8  `ldacs:"name:G_TYPE; size:4; type:set"`
	ASSac uint16 `ldacs:"name:as_sac; size:12; type:set"`
	Sdu   []byte `ldacs:"name:Sdu; type:dbytes"`
}

func ParseGsnfPkt(msg []byte) *GsnfPkt {
	gsnfMsg := GsnfPkt{
		Sdu: make([]byte, len(msg)-GSNF_HEAD_LEN),
	}
	_, err := util.UnmarshalLdacsPkt(msg, &gsnfMsg)
	if err != nil {
		return nil
	}
	return &gsnfMsg
}

func AssembleGsnfPkt(pkt *GsnfPkt) []byte {
	gsnfPdu, err := util.MarshalLdacsPkt(pkt)
	if err != nil {
		global.LOGGER.Error("Failed Assemble Pkt", zap.Error(err))
		return nil
	}

	return gsnfPdu
}
