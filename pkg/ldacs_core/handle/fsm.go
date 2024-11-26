package handle

import (
	"context"
	"errors"
	"github.com/looplab/fsm"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
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
	unit := ctx.Value("unit").(*LdacsUnit)
	st := unit.State

	key, N2, err := GenerateSharedKey(st)
	if err != nil {
		global.LOGGER.Error("Generate Shared key failed.", zap.Error(err))
		return err
	}

	unit.ToSendPkt(&AucResp{
		Stype:  global.AUC_RESP,
		Ver:    st.Ver,
		PID:    global.PID(st.PID),
		ASSac:  uint16(st.AsSac),
		GSSac:  uint16(st.GsSac),
		MacLen: global.MacLen(st.MacLen),
		AuthID: global.AuthID(st.MacLen),
		EncID:  global.EncID(st.MacLen),
		N2:     N2,
		KeyLen: global.KeyLen(st.KdfLen),
	}, key)

	return nil
}
func (s *SecStateFsm) beforeAuthStateG2(ctx context.Context, e *fsm.Event) error {
	return nil
}
func (s *SecStateFsm) afterEvent(ctx context.Context, e *fsm.Event) error {
	unit := ctx.Value("unit").(*LdacsUnit)

	var authSt global.AuthStateKind

	switch e.Dst {
	case global.AUTH_STAGE_G0.GetString():
		authSt = global.AUTH_STAGE_G0
	case global.AUTH_STAGE_G1.GetString():
		authSt = global.AUTH_STAGE_G1
	case global.AUTH_STAGE_G2.GetString():
		authSt = global.AUTH_STAGE_G2
	case global.AUTH_STAGE_UNDEFINED.GetString():
		authSt = global.AUTH_STAGE_UNDEFINED
	default:
		return errors.New("wrong Para")
	}

	if err := unit.TransState(authSt); err != nil {
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
		err := s.FSM.Event(ctx, global.AUTH_STAGE_UNDEFINED.GetString())
		if err != nil {
			return
		}
	}
}

func InitNewAuthFsm() *fsm.FSM {
	return fsm.NewFSM(global.AUTH_STAGE_UNDEFINED.GetString(),
		fsm.Events{
			*getFSMEvents(global.AUTH_STAGE_G0.GetString(), global.AUTH_STAGE_UNDEFINED.GetString()),
			*getFSMEvents(global.AUTH_STAGE_G1.GetString(), global.AUTH_STAGE_G0.GetString()),
			*getFSMEvents(global.AUTH_STAGE_G2.GetString(), global.AUTH_STAGE_G1.GetString()),

			//处理错误
			*getFSMEvents(global.AUTH_STAGE_UNDEFINED.GetString(), global.AUTH_STAGE_G0.GetString(), global.AUTH_STAGE_G1.GetString(), global.AUTH_STAGE_G2.GetString(), global.AUTH_STAGE_UNDEFINED.GetString()),
		},
		fsm.Callbacks{
			"before_" + global.AUTH_STAGE_G0.GetString(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG0(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G1.GetString(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG1(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G2.GetString(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateG2(ctx, e))
			},
			"before_" + global.AUTH_STAGE_UNDEFINED.GetString(): func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.beforeAuthStateUndef(ctx, e))
			},
			"after_event": func(ctx context.Context, e *fsm.Event) {
				SecStates.handleErrEvent(ctx, SecStates.afterEvent(ctx, e))
			},
		},
	)
}
