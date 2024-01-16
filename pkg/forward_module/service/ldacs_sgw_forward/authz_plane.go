package ldacs_sgw_forward

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"

	"gorm.io/gorm"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
)

type AuthzPlaneService struct {
}

// CreateAuthzPlane 创建飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) CreateAuthzPlane(authzPlane *ldacs_sgw_forward.AuthzPlane) (err error) {
	err = global.DB.Create(authzPlane).Error
	return err
}

// DeleteAuthzPlane 删除飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) DeleteAuthzPlane(id string) (err error) {
	err = global.DB.Delete(&ldacs_sgw_forward.AuthzPlane{}, "id = ?", id).Error
	return err
}

// DeleteAuthzPlaneByIds 批量删除飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) DeleteAuthzPlaneByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]ldacs_sgw_forward.AuthzPlane{}, "id in ?", ids).Error
	return err
}

// UpdateAuthzPlane 更新飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) UpdateAuthzPlane(authzPlane ldacs_sgw_forward.AuthzPlane) (err error) {
	err = global.DB.Save(&authzPlane).Error
	return err
}

// GetAuthzPlane 根据id获取飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) GetAuthzPlane(id string) (authzPlane ldacs_sgw_forward.AuthzPlane, err error) {
	err = global.DB.Where("id = ?", id).First(&authzPlane).Error
	return
}

// GetAuthzPlaneInfoList 分页获取飞机业务授权记录
// Author [piexlmax](https://github.com/piexlmax)
func (authzPlaneService *AuthzPlaneService) GetAuthzPlaneInfoList(info ldacs_sgw_forwardReq.AuthzPlaneSearch) (list []ldacs_sgw_forward.AuthzPlane, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&ldacs_sgw_forward.AuthzPlane{})
	var authzPlanes []ldacs_sgw_forward.AuthzPlane
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//if info.Authz_planeId != nil {
	//	db = db.Where("authz_plane_id = ?", info.Authz_planeId)
	//}
	//if info.Authz_flight != nil {
	//	db = db.Where("authz_flight = ?", info.Authz_flight)
	//}
	//if info.Authz_autz != nil {
	//	db = db.Where("authz_autz = ?", info.Authz_autz)
	//}
	//if info.Authz_state != nil {
	//	db = db.Where("authz_state = ?", info.Authz_state)
	//}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	//if limit != 0 {
	//	db = db.Limit(limit).Offset(offset)
	//}

	//err = db.Find(&authzPlanes).Error
	err = db.Joins("Planeid").Joins("Flight").Joins("Authz").Limit(limit).Offset(offset).Find(&authzPlanes).Error
	return authzPlanes, total, err
}
func (authzPlaneService *AuthzPlaneService) GetOptions() (*ldacs_sgw_forward.AuthzOptions, error) {
	db := global.DB
	var authzOpts ldacs_sgw_forward.AuthzOptions
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

func (authzPlaneService *AuthzPlaneService) StateChange(authzPlane *ldacs_sgw_forward.AuthzPlane) (err error) {
	db := global.DB.Model(&ldacs_sgw_forward.AuthzPlane{})
	err = db.Where("id = ?", authzPlane.ID).Update("authz_state", authzPlane.AuthzState).Error

	fmt.Printf("err: %v\n", err)
	return err
}
