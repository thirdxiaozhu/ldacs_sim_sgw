package handle

import (
	"context"
	"encoding/base64"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/backward_module"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"ldacs_sim_sgw/pkg/ldacs_core/service"
	"sync"
	"unsafe"

	"go.uber.org/zap"
)

const GTYPE_LEN = 4
const GSNF_HEAD_LEN = 2
const GSNF_SAC_HEAD_LEN = 4

type LdacsHandler struct {
	ldacsUnits sync.Map //as_sac <-> ld_u_c_node  map
}

type LdacsUnit struct {
	AsSac  uint16 `json:"as_sac"`
	AsUa   uint32
	GsSac  uint16 `json:"gs_sac"`
	ConnID uint32
	State  *model.State
	Fsm    *LdacsStateFsm
	//KUpdateFsm     *LdacsStateFsm // TODO: check this
	HandlerRootKey unsafe.Pointer
	HandlerAsSgw   unsafe.Pointer
	KeyAsGs        []byte
	Nonce          []byte
}

func InitLdacsUnit(connId, asUa uint32, asSac uint16) *LdacsUnit {
	var err error
	unit := &LdacsUnit{
		ConnID: connId,
		AsSac:  asSac,
		AsUa:   asUa,
		GsSac:  0xABD,
		Fsm:    InitNewAuthFsm(),
		//KUpdateFsm: InitNewKUpdateFsm(), // TODO: check this
		KeyAsGs: nil,
		State:   service.InitState(asSac, 10010),
	}

	unit.HandlerRootKey, err = GetKeyHandle("ACTIVE", "ROOT_KEY", 10010, 10000)
	if err != nil {
		global.LOGGER.Error("错误！", zap.Error(err))
		return nil
	}

	// 认证状态机初始化为G0
	ctx := context.Background()
	ctx = context.WithValue(ctx, "unit", unit)
	if err := unit.Fsm.Fsm.Event(ctx, global.AUTH_STAGE_G0.GetString()); err != nil {
		global.LOGGER.Error("错误！", zap.Error(err))
		return nil
	}

	return unit
}

func (u *LdacsUnit) HandleMsg(gsnfSdu []byte) {
	st := u.State
	ctx := context.Background()
	ctx = context.WithValue(ctx, "unit", u)
	//logger.Warn(u.Fsm.Current())
	//for i := range gsnfSdu {
	//	fmt.Printf("%02x ", gsnfSdu[i])
	//}
	//fmt.Println()

	switch global.STYPE(gsnfSdu[0]) {
	case global.AUC_RQST:
		var aucRqst AucRqst

		tail, err := util.UnmarshalLdacsPkt(gsnfSdu, &aucRqst)
		if err != nil {
			return
		}

		isSuccess := VerifyHmac(u.HandlerRootKey, gsnfSdu[:tail], gsnfSdu[tail:], 32)
		if isSuccess == false {
			global.LOGGER.Error("Hmac Verify failed")
			return
		}

		st.Ver = uint8(aucRqst.Ver)
		st.PID = uint8(aucRqst.PID)
		st.MacLen = uint8(aucRqst.MacLen)
		st.AuthId = uint8(aucRqst.AuthID)
		st.EncId = uint8(aucRqst.EncID)

		if err := u.Fsm.Fsm.Event(ctx, global.AUTH_STAGE_G1.GetString()); err != nil {
			return
		}

	case global.AUC_KEY_EXEC:
		var aucKeyExec AucKeyExec

		tail, err := util.UnmarshalLdacsPkt(gsnfSdu, &aucKeyExec)
		if err != nil {
			global.LOGGER.Error("Unmarshel ldacs error", zap.Error(err))
			return
		}

		isSuccess := VerifyHmac(u.HandlerAsSgw, gsnfSdu[:tail], gsnfSdu[tail:], 32)
		if isSuccess == false {
			global.LOGGER.Error("Hmac Verify failed")
			return
		}

		st.Ver = uint8(aucKeyExec.Ver)
		st.PID = uint8(aucKeyExec.PID)
		st.MacLen = uint8(aucKeyExec.MacLen)
		st.AuthId = uint8(aucKeyExec.AuthID)
		st.EncId = uint8(aucKeyExec.EncID)

		if err := u.Fsm.Fsm.Event(ctx, global.AUTH_STAGE_G2.GetString()); err != nil {
			return
		}

		ctxG3 := context.Background()
		ctxG3 = context.WithValue(ctxG3, "unit", u)
		ctxG3 = context.WithValue(ctxG3, "targetGsSAC", uint16(0xABF))
		if err := u.Fsm.Fsm.Event(ctxG3, global.AUTH_STAGE_G3.GetString()); err != nil {
			return
		}
	case global.KUPDATE_RESPONSE:

		var kupdResponse KUpdateResponse
		tail, err := util.UnmarshalLdacsPkt(gsnfSdu, &kupdResponse)
		if err != nil {
			global.LOGGER.Error("Unmarshel ldacs error", zap.Error(err))
			return
		}

		isSuccess := VerifyHmac(u.HandlerAsSgw, gsnfSdu[:tail], gsnfSdu[tail:], 32)
		if isSuccess == false {
			global.LOGGER.Error("Hmac Verify failed")
			return
		}

		if err := u.Fsm.Fsm.Event(ctx, global.AUTH_STAGE_G2.GetString()); err != nil {
			return
		}
	}
}

func (u *LdacsUnit) TransState(newState global.AuthStateKind) error {
	u.State.AuthState = newState
	err := service.AuthcStateSer.NewAuthcStateTrans(
		u.State.AsSac,
		u.State.GsSac,
		u.State.GscSac,
		newState)
	if err != nil {
		return err
	}
	return nil
}

func (u *LdacsUnit) SendPkt(v any, GType GTYPE) {
	sdu, err := util.MarshalLdacsPkt(v)
	if err != nil {
		global.LOGGER.Error("Failed Send", zap.Error(err))
		return
	}

	switch GType {
	case GSNF_SAC_RESP:
		if err = backward_module.SendPkt(AssembleGsnfPkt(&GsnfSacPkt{
			GType: uint8(GType),
			UA:    u.AsUa,
			Sdu:   sdu,
		}), u.ConnID); err != nil {
			global.LOGGER.Error("Failed Send", zap.Error(err))
			return
		}
	case GSNF_SNF_DOWNLOAD, GSNF_GS_KEY_TRANS:
		if GType == GSNF_SNF_DOWNLOAD {
			hmac, _ := util.CalcHMAC(u.HandlerAsSgw, sdu, global.MacLen(u.State.MacLen).GetMacLen())
			sdu = append(sdu, hmac...)
		}
		if err = backward_module.SendPkt(AssembleGsnfPkt(&GsnfPkt{
			GType: uint8(GType),
			ASSac: u.AsSac,
			Sdu:   sdu,
		}), u.ConnID); err != nil {
			global.LOGGER.Error("Failed Send", zap.Error(err))
			return
		}

		if err = service.AuditAsRawSer.NewAuditRaw(u.AsSac, int(global.OriFl), base64.StdEncoding.EncodeToString(sdu)); err != nil {
			return
		}
	default:
		return
	}
}
func (l *LdacsHandler) Serve(msg []byte, connId uint32) {
	gsnfMsg, err := ParseGsnf(msg)
	if err != nil {
		global.LOGGER.Error("Serve Failure", zap.Error(err))
		return
	}

	switch gsnfMsg.(type) {
	case *GsnfPkt:
		gsnfPkt := gsnfMsg.(*GsnfPkt)

		unit, ok := l.ldacsUnits.Load(gsnfPkt.ASSac)
		if ok == false {
			return
		}

		ldacsUnitPtr := unit.(*LdacsUnit)
		ldacsUnitPtr.HandleMsg(gsnfPkt.Sdu)

		if err := service.AuditAsRawSer.NewAuditRaw(gsnfPkt.ASSac, int(global.OriRl), base64.StdEncoding.EncodeToString(gsnfPkt.Sdu)); err != nil {
			return
		}
	case *GsnfSacPkt:
		gsnfSacPkt := gsnfMsg.(*GsnfSacPkt)
		var sac uint16 = 1234

		unit, ok := l.ldacsUnits.Load(sac)
		if ok == false {
			unit = InitLdacsUnit(connId, gsnfSacPkt.UA, sac)
			l.ldacsUnits.Store(sac, unit)
		} else {
			return
		}

		ldacsUnitPtr := unit.(*LdacsUnit)
		ldacsUnitPtr.SendPkt(&GSSacRespSdu{AsSac: sac}, GSNF_SAC_RESP)

	default:
	}

}

func (l *LdacsHandler) Close(id uint32) {

	l.ldacsUnits.Range(func(key, value interface{}) bool {
		asSac := key
		node := value.(*LdacsUnit)
		if node.ConnID == id {
			l.ldacsUnits.Delete(asSac)
		}
		return true
	})
}
