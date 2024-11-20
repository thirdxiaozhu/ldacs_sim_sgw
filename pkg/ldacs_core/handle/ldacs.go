package handle

import (
	"context"
	"encoding/json"
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
	AsSac  uint16 `json:"as_sac"`
	GsSac  uint16 `json:"gs_sac"`
	ConnID uint32
	State  LdacsStateConnNode

	UaGs  uint64          `json:"ua_gs"`
	UaGsc uint64          `json:"ua_gsc"`
	Head  SecHead         `json:"head"`
	Data  json.RawMessage `json:"data"`

	pldA1     SecPldA1
	pldKdf    SecPldKdf
	pldKdfCon SecPldKdfCon
}

func (u *LdacsUnit) HandleMsg(gsnfMsg []byte) {

	var aucRqst AucRqst

	err := util.UnmarshalLdacsPkt(gsnfMsg, &aucRqst)
	if err != nil {
		return
	}

	global.LOGGER.Info("AucRqst Packet", zap.Any("11", aucRqst))
}

const GSNF_HEAD_LEN = 4

type LdacsStateConnNode struct {
	State   *model.State
	SecHead *SecHead
	AuthFsm fsm.FSM
	Conn    *backward_module.GscConn
}

func newUnitNode(unit *LdacsUnit, conn *backward_module.GscConn) *LdacsStateConnNode {
	ctx := context.Background()

	st, err := service.StateSer.FindStateByAsSac(uint64(unit.AsSac))
	if err != nil {
		global.LOGGER.Error("错误！", zap.Error(err))
		return nil
	}

	st.AuthState = global.AUTH_STAGE_G0
	st.GsSac = unit.UaGs
	st.GscSac = unit.UaGsc
	st.SharedKey = util.GetShardKey(uint64(unit.AsSac))

	unitnodeP := &LdacsStateConnNode{
		State:   &st,
		AuthFsm: *InitNewAuthFsm(),
		Conn:    conn,
	}

	ctx = context.WithValue(ctx, "node", unitnodeP)
	if err = unitnodeP.AuthFsm.Event(ctx, global.AUTH_STAGE_G0.GetString()); err != nil {
		global.LOGGER.Error("错误！", zap.Error(err))
		return nil
	}

	return unitnodeP
}

func (node *LdacsStateConnNode) ToSendPkt(unit *LdacsUnit) {
	pktJ, err := json.Marshal(unit)
	if err != nil {
		return
	}

	if err := service.AuditAsRawSer.NewAuditRaw(uint64(unit.AsSac), int(global.OriFl), string(pktJ)); err != nil {
		return
	}

	node.Conn.SendPkt(pktJ)
}

type LdacsHandler struct {
	ldacsUnits sync.Map //as_sac <-> ld_u_c_node  map
}

func (l *LdacsHandler) Serve(msg []byte, connId uint32) {
	global.LOGGER.Info(string(msg), zap.Uint32("ID ", connId))

	gsnfPkt := ParseGsnfPkt(msg)

	unit, _ := l.ldacsUnits.LoadOrStore(gsnfPkt.ASSac, &LdacsUnit{
		ConnID: connId,
		AsSac:  gsnfPkt.ASSac,
		GsSac:  0xABD,
	})

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
	ctx := context.Background()
	st := node.State
	//st.SecHead = unit.Head
	node.SecHead = &unit.Head

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

func TransState(node *LdacsStateConnNode, newState global.AuthStateKind) error {
	node.State.AuthState = newState
	err := service.AuthcStateSer.NewAuthcStateTrans(
		node.State.AsSac,
		node.State.GsSac,
		node.State.GscSac,
		newState)
	if err != nil {
		return err
	}
	return nil
}
