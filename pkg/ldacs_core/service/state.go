package service

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
)

type StateService struct {
}

func (stateService *StateService) CreateState(state *model.State) (err error) {
	err = global.DB.Create(state).Error
	return err
}
