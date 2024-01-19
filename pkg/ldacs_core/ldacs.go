package ldacscore

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/backward_module"
	model "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	service "ldacs_sim_sgw/pkg/forward_module/service/ldacs_sgw_forward"
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
	State *state
	Conn  *backward_module.GscConn
}

func newUnitNode(uas uint32, conn *backward_module.GscConn) *LdacsStateConnNode {

	unitnodeP := &LdacsStateConnNode{
		State: initState(uas),
		Conn:  conn,
	}

	err := unitnodeP.State.AuthFsm.Event(context.Background(), AUTH_STATE_G0.String())
	if err != nil {
		return nil
	}

	return unitnodeP
}

type LdacsHandler struct {
	ldacsConn sync.Map //uas <-> ld_u_c_node  map
}

func MakeLdacsHandler() *LdacsHandler {
	return &LdacsHandler{}
}

func (l *LdacsHandler) ServeGSC(msg []byte, conn *backward_module.GscConn) {

	var unit LdacsUnit
	err := json.Unmarshal(msg, &unit)
	if err != nil {
		return
	}

	accountAsService := service.AccountAsService{}
	if count, err := accountAsService.GetAccountAsBySac(unit.AsSac); count != 0 && err == nil {
		auditService := service.AuditAsRawService{}
		if err := auditService.CreateAuditAsRaw(&model.AuditAsRaw{
			AuditAsSac: int(unit.AsSac),
			AuditAsMsg: string(msg),
		}); err != nil {
			global.LOGGER.Error("失败", zap.Error(err))
		}
	}
	global.LOGGER.Info("成功")

	v, _ := l.ldacsConn.Load(unit.AsSac)
	if v == nil {
		uas := genUAs(unit.AsSac, unit.UaGs, unit.UaGsc)
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
	st.SecHead = unit.Head

	ctx = context.WithValue(ctx, "node", node)
	ctx = context.WithValue(ctx, "unit", unit)
	switch st.SecHead.Cmd {
	case uint8(REGIONAL_ACCESS_REQ):
		err := json.Unmarshal(unit.Data, &unit.pldA1)
		if err != nil {
			return
		}

		st.MacLen = unit.pldA1.MacLen
		st.AuthId = unit.pldA1.AuthID
		st.EncId = unit.pldA1.EncID

		err = st.AuthFsm.Event(ctx, AUTH_STATE_G1.String())
		if err != nil {
			return
		}
	case uint8(REGIONAL_ACCESS_CONFIRM):
		err := json.Unmarshal(unit.Data, &unit.pldKdfCon)
		if err != nil {
			return
		}
		err = st.AuthFsm.Event(ctx, AUTH_STATE_G2.String())
		if err != nil {
			return
		}
	}

}
