package service

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
)

type AuthcStateService struct {
}

// CreateAuthcState 创建认证状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (authcStateService *AuthcStateService) CreateAuthcState(authcState *model.AuthcState) (err error) {
	err = global.DB.Create(authcState).Error
	return err
}

// DeleteAuthcState 删除认证状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (authcStateService *AuthcStateService) DeleteAuthcState(id string) (err error) {
	err = global.DB.Delete(&model.AuthcState{}, "id = ?", id).Error
	return err
}

// DeleteAuthcStateByIds 批量删除认证状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (authcStateService *AuthcStateService) DeleteAuthcStateByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AuthcState{}, "id in ?", ids).Error
	return err
}

// UpdateAuthcState 更新认证状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (authcStateService *AuthcStateService) UpdateAuthcState(authcState model.AuthcState) (err error) {
	err = global.DB.Save(&authcState).Error
	return err
}

// GetAuthcState 根据id获取认证状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (authcStateService *AuthcStateService) GetAuthcState(id string) (authcState model.AuthcState, err error) {
	err = global.DB.Where("id = ?", id).First(&authcState).Error
	return
}

// GetAuthcStateInfoList 分页获取认证状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (authcStateService *AuthcStateService) GetAuthcStateInfoList(info ldacs_sgw_forwardReq.AuthcStateSearch) (list []model.AuthcState, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AuthcState{})
	var authcStates []model.AuthcState
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AuthcAsSac != nil {
		db = db.Where("authc_as_sac = ?", info.AuthcAsSac)
	}
	if info.AuthcGsSac != nil {
		db = db.Where("authc_gs_sac = ?", info.AuthcGsSac)
	}
	if info.AuthcGscSac != nil {
		db = db.Where("authc_gsc_sac = ?", info.AuthcGscSac)
	}
	if info.AuthcState != nil {
		db = db.Where("authc_state = ?", info.AuthcState)
	}
	if info.StartAuthcTransTime != nil && info.EndAuthcTransTime != nil {
		db = db.Where("authc_trans_time BETWEEN ? AND ? ", info.StartAuthcTransTime, info.EndAuthcTransTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&authcStates).Error
	return authcStates, total, err
}
func (authcStateService *AuthcStateService) NewAuthcStateTrans(asSac, gsSac, gscSac uint64, newState global.AuthStateKind) error {
	accountAs, err := AccountAsSer.GetAccountAsBySac(asSac)
	if err != nil {
		global.LOGGER.Error("Failure", zap.Error(err))
		return err
	}
	if err := authcStateService.CreateAuthcState(&model.AuthcState{
		AsSac:       accountAs,
		AuthcGsSac:  gsSac,
		AuthcGscSac: gscSac,
		AuthcState:  newState,
	}); err != nil {
		global.LOGGER.Error("Failure", zap.Error(err))
		return err
	}

	global.LOGGER.Info("成功")
	return nil
}
