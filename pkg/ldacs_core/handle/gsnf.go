package handle

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
)

type GTYPE uint8

const (
	GSNF_SAC_RQST     GTYPE = 0x00
	GSNF_SAC_RESP     GTYPE = 0x01
	GSNF_INITIAL_MSG  GTYPE = 0x02
	GSNF_SNF_UPLOAD   GTYPE = 0x03
	GSNF_SNF_DOWNLOAD GTYPE = 0x04
	GSNF_GS_KEY_TRANS GTYPE = 0x05
)

func (f GTYPE) GetString() string {
	return [...]string{
		"GS_SAC_RQST",
		"GS_SAC_RESP",
		"GS_INITIAL_MSG",
		"GS_SNF_UPLOAD",
		"GS_SNF_DOWNLOAD",
		"GSNF_GS_KEY_TRANS",
	}[f-GSNF_SAC_RQST]
}

func (f GTYPE) CheckValid() bool {
	return f <= GSNF_GS_KEY_TRANS
}

type GsnfSacPkt struct {
	GType uint8  `ldacs:"name:G_TYPE; size:4; type:set"`
	UA    uint32 `ldacs:"name:UA; size:28; type:set"`
	Sdu   []byte `ldacs:"name:Sdu; type:dbytes"`
}

type GsnfPkt struct {
	GType uint8  `ldacs:"name:G_TYPE; size:4; type:set"`
	ASSac uint16 `ldacs:"name:as_sac; size:12; type:set"`
	Sdu   []byte `ldacs:"name:Sdu; type:dbytes"`
}

type GSSacRespSdu struct {
	AsSac uint16 `ldacs:"name:as_sac; size:12; type:set"`
}

type GSNFInitialAsMessage struct {
	GType   GTYPE  `ldacs:"name:GroundTYPE; size:8; type:set"`
	Version uint8  `ldacs:"name:Version; size:4; type:set"`
	ASSac   uint16 `ldacs:"name:as_sac; size:12; type:set"`
	UA      uint32 `ldacs:"name:UA; size:28; type:set"`
	Element uint8  `ldacs:"name:Element Type; size:4; type:set"`
	Sdu     []byte `ldacs:"name:Sdu; type:dbytes"`
}

type GSNFMsgTrans struct {
	GType   GTYPE  `ldacs:"name:GroundTYPE; size:8; type:set"`
	Version uint8  `ldacs:"name:Version; size:4; type:set"`
	ASSac   uint16 `ldacs:"name:as_sac; size:12; type:set"`
	Element uint8  `ldacs:"name:Element Type; size:4; type:set"`
	Sdu     []byte `ldacs:"name:Sdu; type:dbytes"`
}

func ParseGsnf(msg []byte) (any, error) {
	switch global.CONFIG.System.ConnectMode {
	case "GS":
		switch msg[0] >> (util.BITS_PER_BYTE - GTYPE_LEN) & (0xFF >> (util.BITS_PER_BYTE - GTYPE_LEN)) {
		case byte(GSNF_SAC_RQST):
			gsnfSacMsg := GsnfSacPkt{
				Sdu: make([]byte, len(msg)-GSNF_SAC_HEAD_LEN),
			}
			_, err := util.UnmarshalLdacsPkt(msg, &gsnfSacMsg)
			if err != nil {
				return nil, err
			}
			return &gsnfSacMsg, nil
		case byte(GSNF_SNF_DOWNLOAD):
			gsnfMsg := GsnfPkt{
				Sdu: make([]byte, len(msg)-GSNF_HEAD_LEN),
			}
			_, err := util.UnmarshalLdacsPkt(msg, &gsnfMsg)
			if err != nil {
				return nil, err
			}
			return &gsnfMsg, nil
		}
	case "GSC":
		return nil, nil
	default:
		return nil, nil
	}
	return nil, nil
}

func AssembleGsnfPkt(pkt any) []byte {
	gsnfPdu, err := util.MarshalLdacsPkt(pkt)
	if err != nil {
		global.LOGGER.Error("Failed Assemble Pkt", zap.Error(err))
		return nil
	}

	return gsnfPdu
}
