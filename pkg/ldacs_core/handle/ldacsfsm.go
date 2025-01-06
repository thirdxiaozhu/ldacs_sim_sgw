package handle

import (
	"context"
	"errors"
	"github.com/looplab/fsm"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
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

	err := GenerateSharedKey(unit)
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
		N2:     unit.Nonce,
		KeyLen: global.KeyLen(st.KdfLen),
	}, GSNF_CTRL_MSG)

	return nil
}

/* 网关：分发密钥给GS */
func (s *LdacsStateFsm) afterAuthStateG1(ctx context.Context, e *fsm.Event) error {
	unit := ctx.Value("unit").(*LdacsUnit)

	unit.SendPkt(&GSKeyTrans{
		Key:   unit.KeyAsGs,
		Nonce: unit.Nonce,
	}, GSNF_GS_KEY)
	return nil
}

/* 网关：更新主密钥 发送密钥更新请求给AS */
func (s *LdacsStateFsm) beforeAuthStateG3(ctx context.Context, e *fsm.Event) error {
	unit := ctx.Value("unit").(*LdacsUnit)
	st := unit.State

	// generate random N4
	N4 := GenerateRandomBytes(16)

	// update masterkey kas-gs
	err := SGWUpdateMK(util.UAformat(10010), util.UAformat(10000), util.UAformat(10086), util.UAformat(10001), N4) 
	if err != nil {
		global.LOGGER.Error("SGWUpdateMK failed.", zap.Error(err))
		return err
	}
	
	// send kupdate request
	unit.SendPkt(&KUpdateRequest{
		KeyType: global.MASTER_KEY_AS_GS_128,
		SGSSac: util.UAformat(10000), 
		TGSSAC: util.UAformat(10001)
		N4:      N4,         
	}, GSNF_CTRL_MSG)
	return nil
}

/* 网关：分发密钥给GS */
func (s *LdacsStateFsm) afterAuthStateG3(ctx context.Context, e *fsm.Event) error { 
	unit := ctx.Value("unit").(*LdacsUnit)

	// query key value 
	result, err := SGWQueryKeyValueByOwner(util.UAformat(10010), util.UAformat(10000),util.MASTER_KEY_AS_GS, util.ACTIVE);
	if err != nil {
		global.LOGGER.Error("Error querying key-value", zap.Error(err))
	}
	
	unit.SendPkt(&GSKeyTrans{
		KeyType: global.MASTER_KEY_AS_GS_128,
		Key:   result.key,
		N4: unit.Nonce,
	}, GSNF_GS_KEY)

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
	case global.AUTH_STAGE_G3.GetString():
		authSt = global.AUTH_STAGE_G3
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

//func (s *LdacsStateFsm) beforeKUpdateStateG0(ctx context.Context, e *fsm.Event) error {
//	return nil
//}

// // TODO : check - send kupdate request pkt
//
//	func (s *LdacsStateFsm) beforeKUpdateStateG1(ctx context.Context, e *fsm.Event) error {
//		unit := ctx.Value("unit").(*LdacsUnit)
//		st := unit.State
//
//		// generate random N4
//		N4 = GenerateRandomBytes(16)
//
//		unit.SendPkt(&KUpdateRequest{
//			Stype:  global.KUPDATE_REQUEST,
//			Ver:    st.Ver,
//			PID:    global.PID(st.PID),
//			ASSac:  uint16(st.AsSac),
//			KeyType:global.KeyType(st.KeyType),
//			SGSSac: uint16(st.GsSac),
//	       TGSSAC: uint16(st.GsTac),// TODO: check
//			AuthID: global.AuthID(st.AuthId),
//			EncID:  global.EncID(st.EncId),
//			N4:     N4,
//		})
//
//		return nil
//	}
//
// // TODO : check - send kupdate ktransport pkt
//
//	func (s *LdacsStateFsm) beforeKUpdateStateG2(ctx context.Context, e *fsm.Event) error {
//		unit := ctx.Value("unit").(*LdacsUnit)
//		st := unit.State
//
//		// update masterkey kas-gs
//		err := SGWUpdateMK(asUa, gsUa, sgwUa, gstUa, nonce) // TODO: check - parameter source
//		if err != nil {
//			global.LOGGER.Error("SGWUpdateMK failed.", zap.Error(err))
//			return err
//		}
//
//		// generate random N4
//		N4 = GenerateRandomBytes(16)
//
//		unit.SendPkt(&KUpdateRequest{
//			Gtype:  global.KEY_UPDATE_KTRANPORT,
//			Ver:    st.Ver,
//			ASSac:  uint16(st.AsSac),
//			ElementType: global.ELEMENTTYPE(st.ElementType),
//			ElementLength: st.Element_Length // TODO: check
//			KeyType:global.KeyType(st.KeyType),
//			Key:     st.KeyAsGs, // TODO: check
//			N4:     N4, // TODO: storage N4
//		})
//
//		return nil
//	}
//
// /*===================================*/
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
			*getFSMEvents(global.AUTH_STAGE_G2.GetString(), global.AUTH_STAGE_G1.GetString(), global.AUTH_STAGE_G3.GetString()),
			*getFSMEvents(global.AUTH_STAGE_G3.GetString(), global.AUTH_STAGE_G2.GetString()),

			//处理错误
			*getFSMEvents(global.AUTH_STAGE_UNDEFINED.GetString(), global.AUTH_STAGE_G0.GetString(), global.AUTH_STAGE_G1.GetString(), global.AUTH_STAGE_G2.GetString(), global.AUTH_STAGE_G3.GetString(), global.AUTH_STAGE_UNDEFINED.GetString()),
		},
		fsm.Callbacks{
			"before_" + global.AUTH_STAGE_G0.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateG0(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G1.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateG1(ctx, e))
			},
			// 考虑改成after G1， 因为可能不是每一次转为就绪状态都需要给GS发密钥
			"after_" + global.AUTH_STAGE_G1.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.afterAuthStateG1(ctx, e))
			},
			"before_" + global.AUTH_STAGE_G3.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeAuthStateG3(ctx, e))
			},
			"after_" + global.AUTH_STAGE_G3.GetString(): func(ctx context.Context, e *fsm.Event) {
				LdacsFsm.handleErrEvent(ctx, LdacsFsm.afterAuthStateG3(ctx, e))
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

///*===================================*/
//// TODO : check
//func InitNewKUpdateFsm() *LdacsStateFsm {
//	LdacsFsm := &LdacsStateFsm{
//		Name: "LdacsStateFsm",
//	}
//	LdacsFsm.Fsm = fsm.NewFSM(global.KUPDATE_STAGE_UNDEFINED.GetString(),
//		fsm.Events{
//			*getFSMEvents(global.KUPDATE_STAGE_G0.GetString(), global.KUPDATE_STAGE_UNDEFINED.GetString()),
//			*getFSMEvents(global.KUPDATE_STAGE_G1.GetString(), global.KUPDATE_STAGE_G0.GetString()),
//			*getFSMEvents(global.KUPDATE_STAGE_G2.GetString(), global.KUPDATE_STAGE_G1.GetString()),
//
//			//处理错误
//			*getFSMEvents(global.KUPDATE_STAGE_UNDEFINED.GetString(), global.KUPDATE_STAGE_G0.GetString(), global.KUPDATE_STAGE_G1.GetString(), global.KUPDATE_STAGE_G2.GetString(), global.KUPDATE_STAGE_UNDEFINED.GetString()),
//		},
//		fsm.Callbacks{
//			"before_" + global.KUPDATE_STAGE_G0.GetString(): func(ctx context.Context, e *fsm.Event) {
//				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUPDATEStateG0(ctx, e))
//			},
//			"before_" + global.KUPDATE_STAGE_G1.GetString(): func(ctx context.Context, e *fsm.Event) {
//				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUppdateStateG1(ctx, e))
//			},
//			"before_" + global.KUPDATE_STAGE_G2.GetString(): func(ctx context.Context, e *fsm.Event) {
//				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUppdateStateG2(ctx, e))
//			},
//			"before_" + global.KUPDATE_STAGE_UNDEFINED.GetString(): func(ctx context.Context, e *fsm.Event) {
//				LdacsFsm.handleErrEvent(ctx, LdacsFsm.beforeKUppdateStateUndef(ctx, e))
//			},
//			"after_event": func(ctx context.Context, e *fsm.Event) {
//				LdacsFsm.handleErrEvent(ctx, LdacsFsm.afterEvent(ctx, e))
//			},
//		},
//	)
//	return LdacsFsm
//}
