package service

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"gorm.io/gorm"
)

type AccountAsService struct {
}

// CreateAccountAs 创建飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) CreateAccountAs(accountAs *model.AccountAs) (err error) {
	accountAs.AsCurrState = 0
	accountAs.AsSac = uint64(util.GenerateRandomInt(global.SAC_LEN))
	/* 创建空状态 */
	accountAs.State = model.NewState()

	err = global.DB.Create(accountAs).Error
	return err
}

// DeleteAccountAs 删除飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) DeleteAccountAs(id string, userID uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.AccountAs{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&model.AccountAs{}, "id = ?", id).Error; err != nil {
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
		if err := tx.Model(&model.AccountAs{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&model.AccountAs{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAccountAs 更新飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) UpdateAccountAs(accountAs model.AccountAs) (err error) {
	err = global.DB.Save(&accountAs).Error
	return err
}

// GetAccountAs 根据id获取飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) GetAccountAs(id string) (accountAs model.AccountAs, err error) {
	//err = global.DB.Model(&model.AccountAs{}).Preload("State.AccountGs").Preload("State.AccountGsc").Where(fmt.Sprintf("`%v`.`id` = ?", model.AccountAs{}.TableName()), id).Joins("Planeid").Joins("Flight").First(&accountAs).Error
	err = global.DB.Model(&model.AccountAs{}).Where(fmt.Sprintf("`%v`.`id` = ?", model.AccountAs{}.TableName()), id).Joins("Planeid").Joins("Flight").Joins("State").First(&accountAs).Error
	return
}

func (accountAsService *AccountAsService) GetAccountAsBySac(sac uint64) (accountAs model.AccountAs, err error) {
	//err = global.DB.Model(&model.AccountAs{}).Where("as_sac = ?", sac).Unscoped().Count(&count).Error
	err = global.DB.Where("as_sac = ?", sac).First(&accountAs).Error
	return
}

// GetAccountAsInfoList 分页获取飞机站账户记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountAsService *AccountAsService) GetAccountAsInfoList(info ldacs_sgw_forwardReq.AccountAsSearch) (list []model.AccountAs, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AccountAs{})
	var accountAss []model.AccountAs
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

	err = db.Joins("Planeid").Joins("Flight").Joins("State").Limit(limit).Offset(offset).Find(&accountAss).Error
	return accountAss, total, err
}
func (accountAsService *AccountAsService) GetOptions() (*model.AccountAsOptions, error) {
	db := global.DB
	var asOpts model.AccountAsOptions
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

func (accountAsService *AccountAsService) StateChange(accountAs *model.AccountAs) (err error) {
	db := global.DB.Model(&model.AccountAs{})
	err = db.Where("id = ?", accountAs.ID).Update("authz_state", accountAs.AsCurrState).Error

	fmt.Printf("err: %v\n", err)
	return err
}
