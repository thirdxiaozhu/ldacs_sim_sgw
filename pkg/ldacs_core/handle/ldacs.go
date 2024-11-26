package handle

import (
	"context"
	"encoding/json"
	gmssl "github.com/GmSSL/GmSSL-Go"
	"github.com/hdt3213/godis/lib/logger"
	"github.com/looplab/fsm"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/backward_module"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"ldacs_sim_sgw/pkg/ldacs_core/service"
	"sync"
)

type LdacsUnit struct {
	AsSac   uint16 `json:"as_sac"`
	GsSac   uint16 `json:"gs_sac"`
	ConnID  uint32
	State   *model.State
	AuthFsm *fsm.FSM

	UaGs  uint64          `json:"ua_gs"`
	UaGsc uint64          `json:"ua_gsc"`
	Head  SecHead         `json:"head"`
	Data  json.RawMessage `json:"data"`

	pldA1     SecPldA1
	pldKdf    SecPldKdf
	pldKdfCon SecPldKdfCon
}

func InitLdacsUnit(connId uint32, asSac uint16) *LdacsUnit {
	unit := &LdacsUnit{
		ConnID:  connId,
		AsSac:   asSac,
		GsSac:   0xABD,
		AuthFsm: InitNewAuthFsm(),
		State:   service.InitState(asSac),
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

const GSNF_HEAD_LEN = 4

type LdacsStateConnNode struct {
	SecHead *SecHead
	State   *model.State
	AuthFsm fsm.FSM
	Conn    *backward_module.GscConn
}

func (u *LdacsUnit) ToSendPkt(v any, key []byte) {
	pdu, err := util.MarshalLdacsPkt(v)
	if err != nil {
		return
	}

	hmac, err := gmssl.NewSm3Hmac(key)
	if err != nil {
		return
	}
	hmac.Update(pdu)
	mac := hmac.GenerateMac()

	pdu = append(pdu, mac...)

	if err = backward_module.SendPkt(pdu, u.ConnID); err != nil {
		return
	}

	if err = service.AuditAsRawSer.NewAuditRaw(u.AsSac, int(global.OriFl), string(pdu)); err != nil {
		return
	}

}

type LdacsHandler struct {
	ldacsUnits sync.Map //as_sac <-> ld_u_c_node  map
}

func (l *LdacsHandler) Serve(msg []byte, connId uint32) {
	global.LOGGER.Info(string(msg), zap.Uint32("ID ", connId))
	gsnfPkt := ParseGsnfPkt(msg)

	unit, _ := l.ldacsUnits.LoadOrStore(gsnfPkt.ASSac, InitLdacsUnit(connId, gsnfPkt.ASSac))

	//global.LOGGER.Info("GSNF Packet", zap.Any("22", gsnfMsg))

	ldacsUnitPtr := unit.(*LdacsUnit)
	ldacsUnitPtr.HandleMsg(gsnfPkt.Sdu)

	//var unit LdacsUnit
	//err := json.Unmarshal(msg, &unit)
	//if err != nil {
	//	return
	//}
	//
	///* add a new audit raw msg */
	//if err := service.AuditAsRawSer.NewAuditRaw(unit.AsSac, int(global.OriRl), string(msg)); err != nil {
	//	return
	//}
	//
	//v, _ := l.ldacsConn.Load(unit.AsSac)
	//if v == nil {
	//	v = newUnitNode(&unit, conn)
	//	l.ldacsConn.Store(unit.AsSac, v)
	//}
	//
	///* Process new msg */
	//ProcessInputMsg(&unit, v.(*LdacsStateConnNode))
	//
	///* Update new service into database */
	//if err = service.StateSer.UpdateState(v.(*LdacsStateConnNode).State); err != nil {
	//	global.LOGGER.Error("错误！", zap.Error(err))
	//}
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

func UpdateState() {

}

func ProcessInputMsg(unit *LdacsUnit, node *LdacsStateConnNode) {
	st := node.State
	//st.SecHead = unit.Head
	node.SecHead = &unit.Head

	ctx := context.Background()
	ctx = context.WithValue(ctx, "node", node)
	switch node.SecHead.Cmd {
	case uint8(REGIONAL_ACCESS_REQ):
		if err := json.Unmarshal(unit.Data, &unit.pldA1); err != nil {
			return
		}

		st.MacLen = unit.pldA1.MacLen
		st.AuthId = unit.pldA1.AuthID
		st.EncId = unit.pldA1.EncID

		if err := node.AuthFsm.Event(ctx, global.AUTH_STAGE_G1.GetString()); err != nil {
			return
		}

	case uint8(REGIONAL_ACCESS_CONFIRM):
		if err := json.Unmarshal(unit.Data, &unit.pldKdfCon); err != nil {
			return
		}

		st.IsSuccess = unit.pldKdfCon.IsOK

		if err := node.AuthFsm.Event(ctx, global.AUTH_STAGE_G2.GetString()); err != nil {
			return
		}
	}
}
