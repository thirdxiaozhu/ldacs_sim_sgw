package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/pkg/forward_module/global"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
)

type AccountAuthzService struct {
}

// CreateAccountAuthz 创建业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) CreateAccountAuthz(accountAuthz *ldacs_sgw_forward.AccountAuthz) (err error) {
	err = global.GVA_DB.Create(accountAuthz).Error
	return err
}

// DeleteAccountAuthz 删除业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) DeleteAccountAuthz(id string) (err error) {
	err = global.GVA_DB.Delete(&ldacs_sgw_forward.AccountAuthz{}, "id = ?", id).Error
	return err
}

// DeleteAccountAuthzByIds 批量删除业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) DeleteAccountAuthzByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ldacs_sgw_forward.AccountAuthz{}, "id in ?", ids).Error
	return err
}

// UpdateAccountAuthz 更新业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) UpdateAccountAuthz(accountAuthz ldacs_sgw_forward.AccountAuthz) (err error) {
	err = global.GVA_DB.Save(&accountAuthz).Error
	return err
}

// GetAccountAuthz 根据id获取业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) GetAccountAuthz(id string) (accountAuthz ldacs_sgw_forward.AccountAuthz, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&accountAuthz).Error
	return
}

// GetAccountAuthzInfoList 分页获取业务权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAuthzService *AccountAuthzService) GetAccountAuthzInfoList(info ldacs_sgw_forwardReq.AccountAuthzSearch) (list []ldacs_sgw_forward.AccountAuthz, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ldacs_sgw_forward.AccountAuthz{})
	var accountAuthzs []ldacs_sgw_forward.AccountAuthz
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
