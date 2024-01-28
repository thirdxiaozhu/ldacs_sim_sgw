package handle

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/looplab/fsm"
	"ldacs_sim_sgw/internal/global"
	"time"
)

type SecStateFsm struct {
	Name string
	FSM  *fsm.FSM
}

var (
	SecStates SecStateFsm
)

func (s *SecStateFsm) beforeAuthStateG0(ctx context.Context, e *fsm.Event) error {
	return nil
}

func (s *SecStateFsm) beforeAuthStateG1(ctx context.Context, e *fsm.Event) error {
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
		KdfKB:  st.KdfKB,
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
func (s *SecStateFsm) beforeAuthStateG2(ctx context.Context, e *fsm.Event) error {
	return nil
}
func (s *SecStateFsm) afterEvent(ctx context.Context, e *fsm.Event) error {
	node := ctx.Value("node").(*LdacsStateConnNode)

	var authSt global.AuthStateKind

	switch e.Dst {
	case global.AUTH_STAGE_G0.String():
		authSt = global.AUTH_STAGE_G0
	case global.AUTH_STAGE_G1.String():
		authSt = global.AUTH_STAGE_G1
	case global.AUTH_STAGE_G2.String():
		authSt = global.AUTH_STAGE_G2
	case global.AUTH_STAGE_UNDEFINED.String():
		authSt = global.AUTH_STAGE_UNDEFINED
	default:
		return errors.New("wrong Para")
	}

	if err := TransState(node, authSt); err != nil {
		return err
	}
	return nil
}

func (s *SecStateFsm) beforeAuthStateUndef(ctx context.Context, e *fsm.Event) error {

	return nil
}

func getFSMEvents(dst string, src ...string) *fsm.EventDesc {
	return &fsm.EventDesc{
		Name: dst,
		Src:  src,
		Dst:  dst,
	}
}

func (s *SecStateFsm) handleErrEvent(ctx context.Context, err error) {
	if err != nil {
		err := s.FSM.Event(ctx, global.AUTH_STAGE_UNDEFINED.String())
		if err != nil {
			return
		}
	}
}

func InitNewAuthFsm() *fsm.FSM {
	return fsm.NewFSM(global.AUTH_STAGE_UNDEFINED.String(),
		fsm.Events{
			*getFSMEvents(global.AUTH_STAGE_G0.String(), global.AUTH_STAGE_UNDEFINED.String()),
			*getFSMEvents(global.AUTH_STAGE_G1.String(), global.AUTH_STAGE_G0.String()),
			*getFSMEvents(global.AUTH_STAGE_G2.String(), global.AUTH_STAGE_G1.String()),

			//处理错误
			*getFSMEvents(global.AUTH_STAGE_UNDEFINED.String(), global.AUTH_STAGE_G0.String(), global.AUTH_STAGE_G1.String(), global.AUTH_STAGE_G2.String(), global.AUTH_STAGE_UNDEFINED.String()),
		},
		fsm.Callbacks{
			"before_" + global.AUTH_STAGE_G0.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG0(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G1.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG1(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G2.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG2(ctx, e))
			},
			"before_" + global.AUTH_STAGE_UNDEFINED.String(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateUndef(ctx, e))
			},
			"after_event": func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.afterEvent(ctx, e))
			},
		},
	)
}
