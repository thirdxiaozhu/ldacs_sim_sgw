package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/pkg/forward_module/f_global"

	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
)

type AccontFlightService struct {
}

// CreateAccontFlight 创建航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) CreateAccontFlight(accountFlight *ldacs_sgw_forward.AccountFlight) (err error) {
	err = f_global.GVA_DB.Create(accountFlight).Error
	return err
}

// DeleteAccontFlight 删除航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) DeleteAccontFlight(id string) (err error) {
	err = f_global.GVA_DB.Delete(&ldacs_sgw_forward.AccountFlight{}, "id = ?", id).Error
	return err
}

// DeleteAccontFlightByIds 批量删除航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) DeleteAccontFlightByIds(ids []string) (err error) {
	err = f_global.GVA_DB.Delete(&[]ldacs_sgw_forward.AccountFlight{}, "id in ?", ids).Error
	return err
}

// UpdateAccontFlight 更新航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) UpdateAccontFlight(accountFlight ldacs_sgw_forward.AccountFlight) (err error) {
	err = f_global.GVA_DB.Save(&accountFlight).Error
	return err
}

// GetAccontFlight 根据id获取航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) GetAccontFlight(id string) (accountFlight ldacs_sgw_forward.AccountFlight, err error) {
	err = f_global.GVA_DB.Where("id = ?", id).First(&accountFlight).Error
	return
}

// GetAccontFlightInfoList 分页获取航班记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountFlightService *AccontFlightService) GetAccontFlightInfoList(info ldacs_sgw_forwardReq.AccontFlightSearch) (list []ldacs_sgw_forward.AccountFlight, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := f_global.GVA_DB.Model(&ldacs_sgw_forward.AccountFlight{})
	var accountFlights []ldacs_sgw_forward.AccountFlight
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
