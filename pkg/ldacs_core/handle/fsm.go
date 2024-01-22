package handle

import (
	"context"
	"encoding/json"
	"github.com/looplab/fsm"
	"ldacs_sim_sgw/internal/global"
	"time"
)

func (s *SecState) beforeAuthStateG0(ctx context.Context, e *fsm.Event) error {
	return nil
}
func (s *SecState) beforeAuthStateG1(ctx context.Context, e *fsm.Event) error {
	node := ctx.Value("node").(*LdacsStateConnNode)
	st := node.State

	node.SecHead.Cmd = uint8(REGIONAL_ACCESS_RES)

	if err := genSharedInfo(st); err != nil {
		return err
	}

	unitData, _ := json.Marshal(SecPldKdf{
		MacLen: st.MacLen,
		AuthID: st.AuthId,
		EncID:  st.EncId,
		RandV:  st.RandV,
		TimeV:  uint64(time.Now().Unix()),
		KdfK:   st.KdfK,
	})

	node.ToSendPkt(&LdacsUnit{
		AsSac: st.AsSac,
		UaGs:  st.GsSac,
		UaGsc: st.GscSac,
		Head:  *node.SecHead,
		Data:  unitData,
	})

	return nil
}
func (s *SecState) beforeAuthStateG2(ctx context.Context, e *fsm.Event) error {
	return nil
}

func (s *SecState) beforeAuthStateUndef(ctx context.Context, e *fsm.Event) error {

	return nil
}

func getFSMEvents(dst string, src ...string) *fsm.EventDesc {
	return &fsm.EventDesc{
		Name: dst,
		Src:  src,
		Dst:  dst,
	}
}

func (s *SecState) handleErrEvent(ctx context.Context, err error) {
	if err != nil {
		err := s.FSM.Event(ctx, global.AUTH_STATE_UNDEFINED.String())
		if err != nil {
			return
		}
	}
}

func InitNewAuthFsm() *fsm.FSM {
	return fsm.NewFSM(global.AUTH_STATE_UNDEFINED.String(),
		fsm.Events{
			*getFSMEvents(global.AUTH_STATE_G0.String(), global.AUTH_STATE_UNDEFINED.String()),
			*getFSMEvents(global.AUTH_STATE_G1.String(), global.AUTH_STATE_G0.String()),
			*getFSMEvents(global.AUTH_STATE_G2.String(), global.AUTH_STATE_G1.String()),

			//处理错误
			*getFSMEvents(global.AUTH_STATE_UNDEFINED.String(), global.AUTH_STATE_G0.String(), global.AUTH_STATE_G1.String(), global.AUTH_STATE_G2.String(), global.AUTH_STATE_UNDEFINED.String()),
		},
		fsm.Callbacks{
			"before_" + global.AUTH_STATE_G0.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG0(ctx, e))
			},
			"before_" + global.AUTH_STATE_G1.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG1(ctx, e))
			},
			"before_" + global.AUTH_STATE_G2.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG2(ctx, e))
			},
			"before_" + global.AUTH_STATE_UNDEFINED.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateUndef(ctx, e))
			},
		},
	)
}
