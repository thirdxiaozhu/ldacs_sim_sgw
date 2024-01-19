package service

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
)

type AccontFlightService struct {
}

// CreateAccontFlight 创建航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) CreateAccontFlight(accountFlight *model.AccountFlight) (err error) {
	err = global.DB.Create(accountFlight).Error
	return err
}

// DeleteAccontFlight 删除航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) DeleteAccontFlight(id string) (err error) {
	err = global.DB.Delete(&model.AccountFlight{}, "id = ?", id).Error
	return err
}

// DeleteAccontFlightByIds 批量删除航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) DeleteAccontFlightByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AccountFlight{}, "id in ?", ids).Error
	return err
}

// UpdateAccontFlight 更新航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) UpdateAccontFlight(accountFlight model.AccountFlight) (err error) {
	err = global.DB.Save(&accountFlight).Error
	return err
}

// GetAccontFlight 根据id获取航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) GetAccontFlight(id string) (accountFlight model.AccountFlight, err error) {
	err = global.DB.Where("id = ?", id).First(&accountFlight).Error
	return
}

// GetAccontFlightInfoList 分页获取航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) GetAccontFlightInfoList(info ldacs_sgw_forwardReq.AccontFlightSearch) (list []model.AccountFlight, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AccountFlight{})
	var accountFlights []model.AccountFlight
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Flight != "" {
		db = db.Where("flight = ?", info.Flight)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&accountFlights).Error
	return accountFlights, total, err
}
