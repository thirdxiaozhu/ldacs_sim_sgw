package service

import (
	"ldacs_sim_sgw/internal/global"
	request "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"ldacs_sim_sgw/pkg/ldacs_core/model"
)

type AccountGsService struct {
}

// CreateAccountGs 创建地面站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGsService *AccountGsService) CreateAccountGs(accountGs *model.AccountGs) (err error) {
	err = global.DB.Create(accountGs).Error
	return err
}

// DeleteAccountGs 删除地面站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGsService *AccountGsService) DeleteAccountGs(id string) (err error) {
	err = global.DB.Delete(&model.AccountGs{}, "id = ?", id).Error
	return err
}

// DeleteAccountGsByIds 批量删除地面站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGsService *AccountGsService) DeleteAccountGsByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AccountGs{}, "id in ?", ids).Error
	return err
}

// UpdateAccountGs 更新地面站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGsService *AccountGsService) UpdateAccountGs(accountGs model.AccountGs) (err error) {
	err = global.DB.Save(&accountGs).Error
	return err
}

// GetAccountGs 根据id获取地面站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGsService *AccountGsService) GetAccountGs(id string) (accountGs model.AccountGs, err error) {
	err = global.DB.Where("id = ?", id).First(&accountGs).Error
	return
}

// GetAccountGsInfoList 分页获取地面站记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountGsService *AccountGsService) GetAccountGsInfoList(info request.AccountGsSearch) (list []model.AccountGs, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AccountGs{})
	var accountGss []model.AccountGs
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GsSac != nil {
		db = db.Where("gs_sac = ?", info.GsSac)
	}
	if info.LatitudeN != nil {
		db = db.Where("latitude_n = ?", info.LatitudeN)
	}
	if info.LongtitudeE != nil {
		db = db.Where("longtitude_e = ?", info.LongtitudeE)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["latitude_n"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&accountGss).Error
	return accountGss, total, err
}
