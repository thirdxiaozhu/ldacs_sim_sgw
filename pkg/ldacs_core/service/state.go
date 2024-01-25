package service

import (
	"fmt"
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
	fmt.Println("!!!!!!!!!", asSac)
	err = global.DB.Where("as_sac = ?", asSac).First(&state).Error
	return
}

func (stateService *StateService) UpdateState(state *model.State) (err error) {
	err = global.DB.Save(state).Error
	return err
}
