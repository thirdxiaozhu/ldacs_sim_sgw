package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/backward_module"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	"ldacs_sim_sgw/pkg/ldacs_core/service"
	"sync"
)

type LdacsUnit struct {
	AsSac uint8           `json:"as_sac"`
	UaGs  uint8           `json:"ua_gs"`
	UaGsc uint8           `json:"ua_gsc"`
	Head  SecHead         `json:"head"`
	Data  json.RawMessage `json:"data"`

	pldA1     SecPldA1
	pldKdf    SecPldKdf
	pldKdfCon SecPldKdfCon
}

type LdacsStateConnNode struct {
	State   *model.State
	SecHead *SecHead
	Conn    *backward_module.GscConn
}

func InitState(uas uint32) *model.State {
	st := model.State{
		SnpState:  global.SNP_STATE_CONNECTING,
		AuthState: global.AUTH_STATE_G0,
		IsTerm:    false,
		AsSac:     util.ParseUAs(uas, "AS"),
		GsSac:     util.ParseUAs(uas, "GS"),
		GscSac:    util.ParseUAs(uas, "GSC"),
		KdfLen:    19,
		SharedKey: util.GetShardKey(uas),
		AuthFsm:   *InitNewAuthFsm(),
	}

	return &st
}

func newUnitNode(uas uint32, conn *backward_module.GscConn) *LdacsStateConnNode {
	unitnodeP := &LdacsStateConnNode{
		State: InitState(uas),
		Conn:  conn,
	}

	err := unitnodeP.State.AuthFsm.Event(context.Background(), global.AUTH_STATE_G0.String())
	if err != nil {
		return nil
	}

	return unitnodeP
}

func (node *LdacsStateConnNode) ToSendPkt(pktUnit *LdacsUnit) {
	fmt.Println(pktUnit)
	pktJ, _ := json.Marshal(pktUnit)
	node.Conn.SendPkt(pktJ)
}

type LdacsHandler struct {
	ldacsConn sync.Map //uas <-> ld_u_c_node  map
}

func (l *LdacsHandler) ServeGSC(msg []byte, conn *backward_module.GscConn) {

	var unit LdacsUnit
	err := json.Unmarshal(msg, &unit)
	if err != nil {
		return
	}

	/* add a new audit raw msg */
	if err := service.AuditAsRawSer.NewAuditRaw(unit.AsSac, int(global.OriRl), string(msg)); err != nil {
		return
	}

	v, _ := l.ldacsConn.Load(unit.AsSac)
	if v == nil {
		uas := util.GenUAs(unit.AsSac, unit.UaGs, unit.UaGsc)
		v = newUnitNode(uas, conn)
		l.ldacsConn.Store(unit.AsSac, v)
	}

	ProcessMsg(&unit, v.(*LdacsStateConnNode))
}

func (l *LdacsHandler) Close(conn *backward_module.GscConn) {
	l.ldacsConn.Range(func(key, value interface{}) bool {
		uas := key
		node := value.(*LdacsStateConnNode)
		if node.Conn == conn {
			l.ldacsConn.Delete(uas)
		}
		return true
	})
}

func ProcessMsg(unit *LdacsUnit, node *LdacsStateConnNode) {
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

		if err := st.AuthFsm.Event(ctx, global.AUTH_STATE_G1.String()); err != nil {
			return
		}

	case uint8(REGIONAL_ACCESS_CONFIRM):
		if err := json.Unmarshal(unit.Data, &unit.pldKdfCon); err != nil {
			return
		}

		st.IsSuccess = unit.pldKdfCon.IsOK

		if err := st.AuthFsm.Event(ctx, global.AUTH_STATE_G2.String()); err != nil {
			return
		}
	}

}
