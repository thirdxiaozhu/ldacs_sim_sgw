package service

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
)

type AccountAuthzService struct {
}

// CreateAccountAuthz 创建业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) CreateAccountAuthz(accountAuthz *model.AccountAuthz) (err error) {
	err = global.DB.Create(accountAuthz).Error
	return err
}

// DeleteAccountAuthz 删除业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) DeleteAccountAuthz(id string) (err error) {
	err = global.DB.Delete(&model.AccountAuthz{}, "id = ?", id).Error
	return err
}

// DeleteAccountAuthzByIds 批量删除业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) DeleteAccountAuthzByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AccountAuthz{}, "id in ?", ids).Error
	return err
}

// UpdateAccountAuthz 更新业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) UpdateAccountAuthz(accountAuthz model.AccountAuthz) (err error) {
	err = global.DB.Save(&accountAuthz).Error
	return err
}

// GetAccountAuthz 根据id获取业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) GetAccountAuthz(id string) (accountAuthz model.AccountAuthz, err error) {
	err = global.DB.Where("id = ?", id).First(&accountAuthz).Error
	return
}

// GetAccountAuthzInfoList 分页获取业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) GetAccountAuthzInfoList(info ldacs_sgw_forwardReq.AccountAuthzSearch) (list []model.AccountAuthz, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AccountAuthz{})
	var accountAuthzs []model.AccountAuthz
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Authz_name != "" {
		db = db.Where("authz_name = ?", info.Authz_name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&accountAuthzs).Error

	return accountAuthzs, total, err
}
