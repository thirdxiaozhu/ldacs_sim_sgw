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
	ccardBinPath   = "rootkey.bin"
)

type KeyEntityService struct {
}

// CreateKeyEntity 创建密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) CreateKeyEntity(km *model.KeyEntity) (err error) {
	if km.KeyType == "ROOT_KEY" {
		err = util.GenerateRootKey(util.UAformat(km.Owner1), util.UAformat(km.Owner2), km.KeyLen, km.UpdateCycle, global.CONFIG.Sqlite.Dsn(), model.KeyEntity{}.TableName(), defaultBinPath)
		if err != nil {
			global.LOGGER.Error("Can not generate root key:", zap.Error(err))
		}

		err = util.KmWriteFileToCryptocard(defaultBinPath, ccardBinPath)
		if err != nil {
			global.LOGGER.Error("Can not import root key into crypto card:", zap.Error(err))
		}
	}

	new_km, err := kmService.GetKeyEntityByContent(ldacs_sgw_forwardReq.KeyEntitySearch{
		KeyState: "PRE_ACTIVATION",
		KeyType:  km.KeyType,
		Owner1:   util.UAformat(km.Owner1),
		Owner2:   util.UAformat(km.Owner2),
	})
	if err != nil {
		return err
	}

	err = util.EnableKey(global.CONFIG.Sqlite.Dsn(), model.KeyEntity{}.TableName(), new_km.KeyID)
	return nil
}

// DeleteKeyEntity 删除密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) DeleteKeyEntity(id string) (err error) {
	err = global.KeyDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.KeyEntity{}).Where("id = ?", id).Error; err != nil {
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
	err = global.KeyDB.Transaction(func(tx *gorm.DB) error {
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

// GetKeyEntityByID 根据id获取密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) GetKeyEntityByID(id string) (km model.KeyEntity, err error) {
	err = global.KeyDB.Where("id = ?", id).First(&km).Error
	return
}

// GetKeyEntityByID 根据内容获取密钥记录
// Author [piexlmax](https://github.com/piexlmax)
func (kmService *KeyEntityService) GetKeyEntityByContent(info ldacs_sgw_forwardReq.KeyEntitySearch) (km model.KeyEntity, err error) {
	//err = global.KeyDB.Where("id = ?", id).First(&km).Error
	db := global.KeyDB.Model(&model.KeyEntity{})

	if info.KeyID != "" {
		db = db.Where("key_id = ?", info.KeyID)
	}
	if info.KeyType != "" {
		db = db.Where("key_type = ?", info.KeyType)
	}
	if info.Owner1 != "" {
		db = db.Where("owner1 = ?", util.UAformat(info.Owner1))
	}
	if info.Owner2 != "" {
		db = db.Where("owner2 = ?", util.UAformat(info.Owner2))
	}
	if info.KeyState != "" {
		db = db.Where("key_state = ?", info.KeyState)
	}

	err = db.Order("creatime desc").First(&km).Error
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
	if info.KeyType != "" {
		db = db.Where("key_type = ?", info.KeyType)
	}
	if info.Owner1 != "" {
		db = db.Where("owner1 = ?", util.UAformat(info.Owner1))
	}
	if info.Owner2 != "" {
		db = db.Where("owner2 = ?", util.UAformat(info.Owner2))
	}
	if info.KeyState != "" {
		db = db.Where("key_state = ?", info.KeyState)
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
