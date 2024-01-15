package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/pkg/forward_module/forward_global"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
)

type AccountPlaneService struct {
}

// CreateAccountPlane 创建飞机账户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountplaneService *AccountPlaneService) CreateAccountPlane(accountplane *ldacs_sgw_forward.AccountPlane) (err error) {
	err = forward_global.GVA_DB.Create(accountplane).Error
	return err
}

// DeleteAccountPlane 删除飞机账户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountplaneService *AccountPlaneService) DeleteAccountPlane(id string) (err error) {
	err = forward_global.GVA_DB.Delete(&ldacs_sgw_forward.AccountPlane{}, "id = ?", id).Error
	return err
}

// DeleteAccountPlaneByIds 批量删除飞机账户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountplaneService *AccountPlaneService) DeleteAccountPlaneByIds(ids []string) (err error) {
	err = forward_global.GVA_DB.Delete(&[]ldacs_sgw_forward.AccountPlane{}, "id in ?", ids).Error
	return err
}

// UpdateAccountPlane 更新飞机账户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountplaneService *AccountPlaneService) UpdateAccountPlane(accountplane ldacs_sgw_forward.AccountPlane) (err error) {
	err = forward_global.GVA_DB.Save(&accountplane).Error
	return err
}

// GetAccountPlane 根据id获取飞机账户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountplaneService *AccountPlaneService) GetAccountPlane(id string) (accountplane ldacs_sgw_forward.AccountPlane, err error) {
	err = forward_global.GVA_DB.Where("id = ?", id).First(&accountplane).Error
	return
}

// GetAccountPlaneInfoList 分页获取飞机账户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountplaneService *AccountPlaneService) GetAccountPlaneInfoList(info ldacs_sgw_forwardReq.AccountPlaneSearch) (list []ldacs_sgw_forward.AccountPlane, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := forward_global.GVA_DB.Model(&ldacs_sgw_forward.AccountPlane{})
	var accountplanes []ldacs_sgw_forward.AccountPlane
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Plane_id != "" {
		db = db.Where("plane_id = ?", info.Plane_id)
	}
	if info.Company != "" {
		db = db.Where("company = ?", info.Company)
	}
	if info.Model != "" {
		db = db.Where("model = ?", info.Model)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&accountplanes).Error
	return accountplanes, total, err
}
