package service

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
)

const (
	defaultBinPath = "/root/ldacs/ldacs_sim_sgw/resources/keystore/rootkey.bin"
)

type KeyEntityService struct {
}

// CreateKeyEntity 创建密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) CreateKeyEntity(km *model.KeyEntity) (err error) {
	err = util.GenerateRootKey(km.Owner1, km.Owner2, km.KeyLen, km.UpdateCycle, global.CONFIG.Sqlite.Dsn(), model.KeyEntity{}.TableName(), defaultBinPath)
	if err != nil {
		global.LOGGER.Error("Can not generate root key:", zap.Error(err))
	}
	//util.QueryID()
	err = util.EnableKey(global.CONFIG.Sqlite.Dsn(), model.KeyEntity{}.TableName(), "")
	//err = global.DB.Create(km).Error
	return nil
}

// DeleteKeyEntity 删除密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) DeleteKeyEntity(id string, userID uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.KeyEntity{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&model.KeyEntity{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteKeyEntityByIds 批量删除密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) DeleteKeyEntityByIds(ids []string, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.KeyEntity{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&model.KeyEntity{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateKeyEntity 更新密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) UpdateKeyEntity(km model.KeyEntity) (err error) {
	err = global.KeyDB.Save(&km).Error
	return err
}

// GetKeyEntity 根据id获取密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) GetKeyEntity(id string) (km model.KeyEntity, err error) {
	err = global.KeyDB.Where("id = ?", id).First(&km).Error
	return
}

// GetKeyEntityInfoList 分页获取密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) GetKeyEntityInfoList(info ldacs_sgw_forwardReq.KeyEntitySearch) (list []model.KeyEntity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.KeyDB.Model(&model.KeyEntity{})
	var kms []model.KeyEntity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.KeyID != "" {
		db = db.Where("key_id = ?", info.KeyID)
	}
	if info.Kind != "" {
		db = db.Where("kind = ?", info.Kind)
	}
	if info.User1 != nil {
		db = db.Where("user1 = ?", info.User1)
	}
	if info.User2 != nil {
		db = db.Where("user2 = ?", info.User2)
	}
	if info.KeyStatus != "" {
		db = db.Where("key_status = ?", info.KeyStatus)
	}
	if info.Ciphertext != "" {
		db = db.Where("ciphertext = ?", info.Ciphertext)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&kms).Error
	return kms, total, err
}

func (kmService *KeyEntityService) GetOptions() (*model.KeyEntityOptions, error) {
	db := global.DB
	var keyOpts model.KeyEntityOptions
	var find *gorm.DB

	for {
		find = db.Order("ID asc").Find(&keyOpts.AsPlaneIds)
		if find.Error != nil {
			break
		}
		return &keyOpts, find.Error
	}

	fmt.Printf("find查询失败，err：%v\n", find.Error)

	return nil, find.Error
}
