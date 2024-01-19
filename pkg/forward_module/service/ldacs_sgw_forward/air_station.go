package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"

	"gorm.io/gorm"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
)

type AirStationService struct {
}

// CreateAirStation 创建飞机站记录
// Author [piexlmax](https://github.com/piexlmax)
func (airStationService *AirStationService) CreateAirStation(airStation *ldacs_sgw_forward.AirStation) (err error) {
	err = global.DB.Create(airStation).Error
	return err
}

// DeleteAirStation 删除飞机站记录
// Author [piexlmax](https://github.com/piexlmax)
func (airStationService *AirStationService) DeleteAirStation(id string, userID uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ldacs_sgw_forward.AirStation{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ldacs_sgw_forward.AirStation{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAirStationByIds 批量删除飞机站记录
// Author [piexlmax](https://github.com/piexlmax)
func (airStationService *AirStationService) DeleteAirStationByIds(ids []string, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ldacs_sgw_forward.AirStation{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&ldacs_sgw_forward.AirStation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAirStation 更新飞机站记录
// Author [piexlmax](https://github.com/piexlmax)
func (airStationService *AirStationService) UpdateAirStation(airStation ldacs_sgw_forward.AirStation) (err error) {
	err = global.DB.Save(&airStation).Error
	return err
}

// GetAirStation 根据id获取飞机站记录
// Author [piexlmax](https://github.com/piexlmax)
func (airStationService *AirStationService) GetAirStation(id string) (airStation ldacs_sgw_forward.AirStation, err error) {
	err = global.DB.Where("id = ?", id).First(&airStation).Error
	return
}

// GetAirStationInfoList 分页获取飞机站记录
// Author [piexlmax](https://github.com/piexlmax)
func (airStationService *AirStationService) GetAirStationInfoList(info ldacs_sgw_forwardReq.AirStationSearch) (list []ldacs_sgw_forward.AirStation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&ldacs_sgw_forward.AirStation{})
	var airStations []ldacs_sgw_forward.AirStation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AsPlaneId != nil {
		db = db.Where("as_plane_id = ?", info.AsPlaneId)
	}
	if info.AsFlight != nil {
		db = db.Where("as_flight = ?", info.AsFlight)
	}
	if info.FlightDate != nil {
		db = db.Where("flight_date = ?", info.FlightDate)
	}
	if info.AsSac != "" {
		db = db.Where("as_sac = ?", info.AsSac)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&airStations).Error
	return airStations, total, err
}
