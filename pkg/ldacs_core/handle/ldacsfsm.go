package handle

import (
	"context"
	"errors"
	"github.com/looplab/fsm"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
)

type LdacsStateFsm struct {
	Name string
	Fsm  *fsm.FSM
}

func (s *LdacsStateFsm) beforeAuthStateG0(ctx context.Context, e *fsm.Event) error {
	return nil
}

func (s *LdacsStateFsm) beforeAuthStateG1(ctx context.Context, e *fsm.Event) error {
	unit := ctx.Value("unit").(*LdacsUnit)
	st := unit.State

	N2, err := GenerateSharedKey(unit)
	if err != nil {
		global.LOGGER.Error("Generate Shared key failed.", zap.Error(err))
		return err
	}

	unit.SendPkt(&AucResp{
		Stype:  global.AUC_RESP,
		Ver:    st.Ver,
		PID:    global.PID(st.PID),
		ASSac:  uint16(st.AsSac),
		GSSac:  uint16(st.GsSac),
		MacLen: global.MacLen(st.MacLen),
		AuthID: global.AuthID(st.AuthId),
		EncID:  global.EncID(st.EncId),
		N2:     N2,
		KeyLen: global.KeyLen(st.KdfLen),
	})

	return nil
}
func (s *LdacsStateFsm) beforeAuthStateG2(ctx context.Context, e *fsm.Event) error {
	return nil
}
func (s *LdacsStateFsm) afterEvent(ctx context.Context, e *fsm.Event) error {
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

func (s *LdacsStateFsm) beforeAuthStateUndef(ctx context.Context, e *fsm.Event) error {

	return nil
}

func getFSMEvents(dst string, src ...string) *fsm.EventDesc {
	return &fsm.EventDesc{
		Name: dst,
		Src:  src,
		Dst:  dst,
	}
}

func (s *LdacsStateFsm) handleErrEvent(ctx context.Context, err error) {
	if err != nil {
		err := s.Fsm.Event(ctx, global.AUTH_STAGE_UNDEFINED.GetString())
		if err != nil {
			return
		}
	}
}

func InitNewAuthFsm() *LdacsStateFsm {
	LdacsFsm := &LdacsStateFsm{
		Name: "LdacsStateFsm",
	}
	LdacsFsm.Fsm = fsm.NewFSM(global.AUTH_STAGE_UNDEFINED.GetString(),
		fsm.Events{
			*getFSMEvents(global.AUTH_STAGE_G0.GetString(), global.AUTH_STAGE_UNDEFINED.GetString()),
			*getFSMEvents(global.AUTH_STAGE_G1.GetString(), global.AUTH_STAGE_G0.GetString()),
			*getFSMEvents(global.AUTH_STAGE_G2.GetString(), global.AUTH_STAGE_G1.GetString()),

			//处理错误
			*getFSMEvents(global.AUTH_STAGE_UNDEFINED.GetString(), global.AUTH_STAGE_G0.GetString(), global.AUTH_STAGE_G1.GetString(), global.AUTH_STAGE_G2.GetString(), global.AUTH_STAGE_UNDEFINED.GetString()),
		},
		fsm.Callbacks{
			"before_" + global.AUTH_STAGE_G0.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateG0(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G1.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateG1(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G2.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateG2(ctx, e))
			},
			"before_" + global.AUTH_STAGE_UNDEFINED.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateUndef(ctx, e))
			},
			"after_event": func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.afterEvent(ctx, e))
			},
		},
	)
	return LdacsFsm
}
