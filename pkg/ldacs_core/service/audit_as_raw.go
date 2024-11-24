package service

import (
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"
)

type AuditAsRawService struct {
}

// CreateAuditAsRaw 创建AS报文记录
// Author [piexlmax](https://github.com/piexlmax)
func (auditAsRawService *AuditAsRawService) CreateAuditAsRaw(auditAsRaw *model.AuditAsRaw) (err error) {
	err = global.DB.Create(auditAsRaw).Error
	return err
}

// DeleteAuditAsRaw 删除AS报文记录
// Author [piexlmax](https://github.com/piexlmax)
func (auditAsRawService *AuditAsRawService) DeleteAuditAsRaw(id string) (err error) {
	err = global.DB.Delete(&model.AuditAsRaw{}, "id = ?", id).Error
	return err
}

// DeleteAuditAsRawByIds 批量删除AS报文记录
// Author [piexlmax](https://github.com/piexlmax)
func (auditAsRawService *AuditAsRawService) DeleteAuditAsRawByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]model.AuditAsRaw{}, "id in ?", ids).Error
	return err
}

// UpdateAuditAsRaw 更新AS报文记录
// Author [piexlmax](https://github.com/piexlmax)
func (auditAsRawService *AuditAsRawService) UpdateAuditAsRaw(auditAsRaw model.AuditAsRaw) (err error) {
	err = global.DB.Save(&auditAsRaw).Error
	return err
}

// GetAuditAsRaw 根据id获取AS报文记录
// Author [piexlmax](https://github.com/piexlmax)
func (auditAsRawService *AuditAsRawService) GetAuditAsRaw(id string) (auditAsRaw model.AuditAsRaw, err error) {
	err = global.DB.Where("id = ?", id).First(&auditAsRaw).Error
	return
}

// GetAuditAsRawInfoList 分页获取AS报文记录
// Author [piexlmax](https://github.com/piexlmax)
func (auditAsRawService *AuditAsRawService) GetAuditAsRawInfoList(info ldacs_sgw_forwardReq.AuditAsRawSearch) (list []model.AuditAsRaw, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AuditAsRaw{})
	var auditAsRaws []model.AuditAsRaw
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AuditAsSac != nil {
		db = db.Where("audit_as_sac = ?", info.AuditAsSac)
	}
	if info.AuditAsMsg != "" {
		db = db.Where("audit_as_msg = ?", info.AuditAsMsg)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("AsSac.State").Find(&auditAsRaws).Error
	return auditAsRaws, total, err
}

func (auditAsRawService *AuditAsRawService) NewAuditRaw(asSac uint16, ori int, msg string) error {
	accountAs, err := AccountAsSer.GetAccountAsBySac(asSac)
	if err != nil {
		global.LOGGER.Error("Failure", zap.Error(err))
		return err
	}
	if err := auditAsRawService.CreateAuditAsRaw(&model.AuditAsRaw{
		AsSac:        accountAs,
		AuditLinkOri: ori,
		AuditAsMsg:   msg,
	}); err != nil {
		global.LOGGER.Error("Failure", zap.Error(err))
		return err
	}

	global.LOGGER.Info("成功")
	return nil
}
