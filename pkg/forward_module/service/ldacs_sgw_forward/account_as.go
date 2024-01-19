package ldacs_sgw_forward

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"

	"gorm.io/gorm"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
)

type AccountAsService struct {
}

// CreateAccountAs 创建飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) CreateAccountAs(accountAs *ldacs_sgw_forward.AccountAs) (err error) {
	accountAs.AsState = 0
	accountAs.AsSac = util.GenerateRandomInt(global.SAC_LEN)
	err = global.DB.Create(accountAs).Error
	return err
}

// DeleteAccountAs 删除飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) DeleteAccountAs(id string, userID uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ldacs_sgw_forward.AccountAs{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ldacs_sgw_forward.AccountAs{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAccountAsByIds 批量删除飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) DeleteAccountAsByIds(ids []string, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ldacs_sgw_forward.AccountAs{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&ldacs_sgw_forward.AccountAs{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAccountAs 更新飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) UpdateAccountAs(accountAs ldacs_sgw_forward.AccountAs) (err error) {
	err = global.DB.Save(&accountAs).Error
	return err
}

// GetAccountAs 根据id获取飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) GetAccountAs(id string) (accountAs ldacs_sgw_forward.AccountAs, err error) {
	err = global.DB.Where("id = ?", id).First(&accountAs).Error
	return
}

func (accountAsService *AccountAsService) GetAccountAsBySac(sac uint8) (count int64, err error) {
	err = global.DB.Model(&ldacs_sgw_forward.AccountAs{}).Where("as_sac = ?", sac).Unscoped().Count(&count).Error
	return
}

// GetAccountAsInfoList 分页获取飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) GetAccountAsInfoList(info ldacs_sgw_forwardReq.AccountAsSearch) (list []ldacs_sgw_forward.AccountAs, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&ldacs_sgw_forward.AccountAs{})
	var accountAss []ldacs_sgw_forward.AccountAs
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
	if info.StartAsDate != nil && info.EndAsDate != nil {
		db = db.Where("as_date BETWEEN ? AND ? ", info.StartAsDate, info.EndAsDate)
	}
	if info.AsSac != "" {
		db = db.Where("as_sac = ?", info.AsSac)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	//if limit != 0 {
	//	db = db.Limit(limit).Offset(offset)
	//}

	//err = db.Find(&accountAss).Error

	err = db.Joins("Planeid").Joins("Flight").Limit(limit).Offset(offset).Find(&accountAss).Error
	return accountAss, total, err
}
func (accountAsService *AccountAsService) GetOptions() (*ldacs_sgw_forward.AccountAsOptions, error) {
	db := global.DB
	var asOpts ldacs_sgw_forward.AccountAsOptions
	var find *gorm.DB

	for {
		find = db.Order("ID asc").Find(&asOpts.AsPlaneIds)
		if find.Error != nil {
			break
		}
		find = db.Order("ID asc").Find(&asOpts.AsFlights)
		if find.Error != nil {
			break
		}
		return &asOpts, find.Error
	}

	fmt.Printf("find查询失败，err：%v\n", find.Error)

	return nil, find.Error
}

func (accountAsService *AccountAsService) StateChange(accountAs *ldacs_sgw_forward.AccountAs) (err error) {
	db := global.DB.Model(&ldacs_sgw_forward.AccountAs{})
	err = db.Where("id = ?", accountAs.ID).Update("authz_state", accountAs.AsState).Error

	fmt.Printf("err: %v\n", err)
	return err
}
