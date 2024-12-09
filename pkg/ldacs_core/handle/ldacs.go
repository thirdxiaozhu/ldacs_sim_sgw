package handle

import (
	"context"
	"encoding/base64"
	"github.com/hdt3213/godis/lib/logger"
	"github.com/looplab/fsm"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/backward_module"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"ldacs_sim_sgw/pkg/ldacs_core/service"
	"sync"
	"unsafe"
)

const GSNF_HEAD_LEN = 4

type LdacsHandler struct {
	ldacsUnits sync.Map //as_sac <-> ld_u_c_node  map
}

type LdacsUnit struct {
	AsSac        uint16 `json:"as_sac"`
	GsSac        uint16 `json:"gs_sac"`
	ConnID       uint32
	State        *model.State
	AuthFsm      *fsm.FSM
	HandlerAsSgw unsafe.Pointer
	KeyAsGs      []byte
}

func InitLdacsUnit(connId uint32, asSac uint16) *LdacsUnit {
	unit := &LdacsUnit{
		ConnID:  connId,
		AsSac:   asSac,
		GsSac:   0xABD,
		AuthFsm: InitNewAuthFsm(),
		State:   service.InitState(asSac, 10010),
	}

	//初始化为G0
	ctx := context.Background()
	ctx = context.WithValue(ctx, "unit", unit)
	if err := unit.AuthFsm.Event(ctx, global.AUTH_STAGE_G0.GetString()); err != nil {
		global.LOGGER.Error("错误！", zap.Error(err))
		return nil
	}
	return unit
}

func (u *LdacsUnit) HandleMsg(gsnfMsg []byte) {

	ctx := context.Background()
	st := u.State

	ctx = context.WithValue(ctx, "unit", u)
	logger.Warn(u.AuthFsm.Current())

	switch global.STYPE(gsnfMsg[0]) {
	case global.AUC_RQST:
		var aucRqst AucRqst

		err := util.UnmarshalLdacsPkt(gsnfMsg, &aucRqst)
		if err != nil {
			return
		}

		global.LOGGER.Info("AucRqst Packet", zap.Any("11", aucRqst))

		st.Ver = uint8(aucRqst.Ver)
		st.PID = uint8(aucRqst.PID)
		st.MacLen = uint8(aucRqst.MacLen)
		st.AuthId = uint8(aucRqst.AuthID)
		st.EncId = uint8(aucRqst.EncID)

		if err := u.AuthFsm.Event(ctx, global.AUTH_STAGE_G1.GetString()); err != nil {
			return
		}

		logger.Warn(u.AuthFsm.Current())

	case global.AUC_RESP:
	case global.AUC_KEY_EXEC:

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

func (u *LdacsUnit) SendPkt(v any) {
	sdu, err := util.MarshalLdacsPkt(v)
	if err != nil {
		global.LOGGER.Error("Failed Send", zap.Error(err))
		return
	}

	//hmac, err := gmssl.NewSm3Hmac(key)
	//if err != nil {
	//	global.LOGGER.Error("Failed Send", zap.Error(err))
	//	return
	//}
	//hmac.Update(sdu)
	hmac, err := util.CalcHMAC(u.HandlerAsSgw, sdu, global.MacLen(u.State.MacLen).GetMacLen())
	sdu = append(sdu, hmac...)

	gsnfMsg := GsnfPkt{
		GType:   0,
		Version: 1,
		ASSac:   u.AsSac,
		EleType: 0,
		Sdu:     sdu,
	}

	gsnfPdu, err := util.MarshalLdacsPkt(gsnfMsg)
	if err != nil {
		global.LOGGER.Error("Failed Send", zap.Error(err))
		return
	}

	if err = backward_module.SendPkt(gsnfPdu, u.ConnID); err != nil {
		global.LOGGER.Error("Failed Send", zap.Error(err))
		return
	}

	if err = service.AuditAsRawSer.NewAuditRaw(u.AsSac, int(global.OriFl), base64.StdEncoding.EncodeToString(sdu)); err != nil {
		return
	}

}
func (l *LdacsHandler) Serve(msg []byte, connId uint32) {
	global.LOGGER.Info(string(msg), zap.Uint32("ID ", connId))
	gsnfPkt := ParseGsnfPkt(msg)

	unit, _ := l.ldacsUnits.LoadOrStore(gsnfPkt.ASSac, InitLdacsUnit(connId, gsnfPkt.ASSac))

	ldacsUnitPtr := unit.(*LdacsUnit)
	ldacsUnitPtr.HandleMsg(gsnfPkt.Sdu)
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
