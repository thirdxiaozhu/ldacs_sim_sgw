package service

import (
	"github.com/hdt3213/godis/lib/logger"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
)

type StateService struct {
}

func (stateService *StateService) CreateState(state *model.State) (err error) {
	err = global.DB.Create(state).Error
	return err
}

func (stateService *StateService) FindStateByAsSac(asSac uint64) (state model.State, err error) {
	err = global.DB.Where("as_sac = ?", asSac).First(&state).Error
	return
}

func (stateService *StateService) UpdateState(state *model.State) (err error) {
	err = global.DB.Save(state).Error
	return err
}

func InitState(sac uint16) *model.State {

	//未来需要根据SAC找对应的UA
	s1, err := AccountAsSer.GetAvialAccountAsByUA(654321)

	logger.Warn(s1)
	if err != nil {
		global.LOGGER.Error("错误：%s", zap.Error(err))
		return nil
	}

	s := s1.State
	s.AuthState = global.AUTH_STAGE_G0
	s.GsSac = 0xABD
	s.GscSac = 0xABC

	return s
}
