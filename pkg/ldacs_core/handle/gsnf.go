package handle

import "ldacs_sim_sgw/internal/util"

type GsnfPkt struct {
	GType uint8  `ldacs:"name:G_TYPE; size:8; type:set"`
	ASSac uint16 `ldacs:"name:as_sac; size:12; type:set"`
	Sdu   []byte `ldacs:"name:Sdu; type:dbytes"`
}

func ParseGsnfPkt(msg []byte) *GsnfPkt {
	gsnfMsg := GsnfPkt{
		Sdu: make([]byte, len(msg)-GSNF_HEAD_LEN),
	}
	err := util.UnmarshalLdacsPkt(msg, &gsnfMsg)
	if err != nil {
		return nil
	}
	return &gsnfMsg
}
