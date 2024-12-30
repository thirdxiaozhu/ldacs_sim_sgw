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
/*===================================*/

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
/*===================================*/


func (s *LdacsStateFsm) beforeKUpdateStateG0(ctx context.Context, e *fsm.Event) error {
	return nil
}

// TODO : check - send kupdate request pkt
func (s *LdacsStateFsm) beforeKUpdateStateG1(ctx context.Context, e *fsm.Event) error {
	unit := ctx.Value("unit").(*LdacsUnit)
	st := unit.State

	// generate random N4
	N4 = GenerateRandomBytes(16)

	unit.SendPkt(&KUpdateRequest{
		Stype:  global.KUPDATE_REQUEST,
		Ver:    st.Ver,
		PID:    global.PID(st.PID),
		ASSac:  uint16(st.AsSac),
		KeyType:global.KeyType(st.KeyType),
		SGSSac: uint16(st.GsSac), 
        TGSSAC: uint16(st.GsTac),// TODO: check
		AuthID: global.AuthID(st.AuthId),
		EncID:  global.EncID(st.EncId),
		N4:     N4,
	})

	return nil
}

// TODO : check - send kupdate ktransport pkt
func (s *LdacsStateFsm) beforeKUpdateStateG2(ctx context.Context, e *fsm.Event) error {
	unit := ctx.Value("unit").(*LdacsUnit)
	st := unit.State

	// update masterkey kas-gs
	err := SGWUpdateMK(asUa, gsUa, sgwUa, gstUa, nonce) // TODO: check - parameter source
	if err != nil {
		global.LOGGER.Error("SGWUpdateMK failed.", zap.Error(err))
		return err
	}

	// generate random N4
	N4 = GenerateRandomBytes(16)

	unit.SendPkt(&KUpdateRequest{
		Gtype:  global.KEY_UPDATE_KTRANPORT,
		Ver:    st.Ver,
		ASSac:  uint16(st.AsSac),
		ElementType: global.ELEMENTTYPE(st.ElementType),
		ElementLength: st.Element_Length // TODO: check
		KeyType:global.KeyType(st.KeyType),
		Key:     st.KeyAsGs, // TODO: check
		N4:     N4, // TODO: storage N4
	})

	return nil
}

/*===================================*/
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

/*===================================*/

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

/*===================================*/
// TODO : check
func InitNewKUpdateFsm() *LdacsStateFsm {
	LdacsFsm := &LdacsStateFsm{
		Name: "LdacsStateFsm",
	}
	LdacsFsm.Fsm = fsm.NewFSM(global.KUPDATE_STAGE_UNDEFINED.GetString(),
		fsm.Events{
			*getFSMEvents(global.KUPDATE_STAGE_G0.GetString(), global.KUPDATE_STAGE_UNDEFINED.GetString()),
			*getFSMEvents(global.KUPDATE_STAGE_G1.GetString(), global.KUPDATE_STAGE_G0.GetString()),
			*getFSMEvents(global.KUPDATE_STAGE_G2.GetString(), global.KUPDATE_STAGE_G1.GetString()),

			//处理错误
			*getFSMEvents(global.KUPDATE_STAGE_UNDEFINED.GetString(), global.KUPDATE_STAGE_G0.GetString(), global.KUPDATE_STAGE_G1.GetString(), global.KUPDATE_STAGE_G2.GetString(), global.KUPDATE_STAGE_UNDEFINED.GetString()),
		},
		fsm.Callbacks{
			"before_" + global.KUPDATE_STAGE_G0.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUPDATEStateG0(ctx, e))
			},
			"before_" + global.KUPDATE_STAGE_G1.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUppdateStateG1(ctx, e))
			},
			"before_" + global.KUPDATE_STAGE_G2.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUppdateStateG2(ctx, e))
			},
			"before_" + global.KUPDATE_STAGE_UNDEFINED.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUppdateStateUndef(ctx, e))
			},
			"after_event": func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.afterEvent(ctx, e))
			},
		},
	)
	return LdacsFsm
}