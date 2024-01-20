package service

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
)

type AccountGscService struct {
}

// CreateAccountGsc 创建地面控制站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGscService *AccountGscService) CreateAccountGsc(accountGsc *model.AccountGsc) (err error) {
	err = global.DB.Create(accountGsc).Error
	return err
}

// DeleteAccountGsc 删除地面控制站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGscService *AccountGscService) DeleteAccountGsc(id string) (err error) {
	err = global.DB.Delete(&model.AccountGsc{}, "id = ?", id).Error
	return err
}

// DeleteAccountGscByIds 批量删除地面控制站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGscService *AccountGscService) DeleteAccountGscByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AccountGsc{}, "id in ?", ids).Error
	return err
}

// UpdateAccountGsc 更新地面控制站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGscService *AccountGscService) UpdateAccountGsc(accountGsc model.AccountGsc) (err error) {
	err = global.DB.Save(&accountGsc).Error
	return err
}

// GetAccountGsc 根据id获取地面控制站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGscService *AccountGscService) GetAccountGsc(id string) (accountGsc model.AccountGsc, err error) {
	err = global.DB.Where("id = ?", id).First(&accountGsc).Error
	return
}

// GetAccountGscInfoList 分页获取地面控制站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGscService *AccountGscService) GetAccountGscInfoList(info ldacs_sgw_forwardReq.AccountGscSearch) (list []model.AccountGsc, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AccountGsc{})
	var accountGscs []model.AccountGsc
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GscSac != nil {
		db = db.Where("gsc_sac = ?", info.GscSac)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&accountGscs).Error
	return accountGscs, total, err
}
