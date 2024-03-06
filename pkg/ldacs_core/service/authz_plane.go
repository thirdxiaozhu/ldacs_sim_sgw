package service

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"gorm.io/gorm"
)

type AuthzPlaneService struct {
}

// CreateAuthzPlane 创建飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) CreateAuthzPlane(authzPlane *model.AuthzPlane) (err error) {
	authzPlane.AuthzState = 0
	err = global.DB.Create(authzPlane).Error
	return err
}

// DeleteAuthzPlane 删除飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) DeleteAuthzPlane(id string) (err error) {
	err = global.DB.Delete(&model.AuthzPlane{}, "id = ?", id).Error
	return err
}

// DeleteAuthzPlaneByIds 批量删除飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) DeleteAuthzPlaneByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AuthzPlane{}, "id in ?", ids).Error
	return err
}

// UpdateAuthzPlane 更新飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) UpdateAuthzPlane(authzPlane model.AuthzPlane) (err error) {
	err = global.DB.Save(&authzPlane).Error
	return err
}

// GetAuthzPlane 根据id获取飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) GetAuthzPlane(id string) (authzPlane model.AuthzPlane, err error) {
	err = global.DB.Preload("AccountAs.State").Joins("Authz").Where(fmt.Sprintf("`%v`.`id` = ?", model.AuthzPlane{}.TableName()), id).First(&authzPlane).Error
	return
}

// GetAuthzPlaneInfoList 分页获取飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) GetAuthzPlaneInfoList(info ldacs_sgw_forwardReq.AuthzPlaneSearch) (list []model.AuthzPlane, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AuthzPlane{})
	var authzPlanes []model.AuthzPlane
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//if info.Authz_planeId != nil {
	//	db = db.Where("authz_plane_id = ?", info.Authz_planeId)
	//}
	//if info.Authz_autz != nil {
	//	db = db.Where("authz_autz = ?", info.Authz_autz)
	//}
	if info.Authz_as != nil {
		db = db.Where("authz_as = ?", info.Authz_as)
	}
	if info.Authz_state != nil {
		db = db.Where("authz_state = ?", info.Authz_state)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	//if limit != 0 {
	//	db = db.Limit(limit).Offset(offset)
	//}

	//err = db.Find(&authzPlanes).Error
	err = db.Preload("AccountAs.State").Joins("Authz").Limit(limit).Offset(offset).Find(&authzPlanes).Error
	return authzPlanes, total, err
}

func (authzPlaneService *AuthzPlaneService) GetOptions() (*model.AuthzOptions, error) {
	db := global.DB
	var authzOpts model.AuthzOptions
	var find *gorm.DB

	for {
		find = db.Order("ID asc").Find(&authzOpts.AuthzAuthzs)
		if find.Error != nil {
			break
		}
		find = db.Order("ID asc").Find(&authzOpts.AuthzPlaneIds)
		if find.Error != nil {
			break
		}
		find = db.Order("ID asc").Find(&authzOpts.AuthzFlights)
		if find.Error != nil {
			break
		}
		return &authzOpts, find.Error
	}

	fmt.Printf("find查询失败，err：%v\n", find.Error)

	return nil, find.Error
}

func (authzPlaneService *AuthzPlaneService) StateChange(authzPlane *model.AuthzPlane) (err error) {
	db := global.DB.Model(&model.AuthzPlane{})
	err = db.Where("id = ?", authzPlane.ID).Update("authz_state", authzPlane.AuthzState).Error

	fmt.Printf("err: %v\n", err)
	return err
}
