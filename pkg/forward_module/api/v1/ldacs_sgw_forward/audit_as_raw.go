package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/service"
)

type AuditAsRawApi struct {
}

var auditAsRawService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AuditAsRawService

// CreateAuditAsRaw 创建AS报文
// @Tags AuditAsRaw
// @Summary 创建AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuditAsRaw true "创建AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /auditAsRaw/createAuditAsRaw [post]
func (auditAsRawApi *AuditAsRawApi) CreateAuditAsRaw(c *gin.Context) {
	var auditAsRaw model.AuditAsRaw
	err := c.ShouldBindJSON(&auditAsRaw)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := auditAsRawService.CreateAuditAsRaw(&auditAsRaw); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAuditAsRaw 删除AS报文
// @Tags AuditAsRaw
// @Summary 删除AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuditAsRaw true "删除AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /auditAsRaw/deleteAuditAsRaw [delete]
func (auditAsRawApi *AuditAsRawApi) DeleteAuditAsRaw(c *gin.Context) {
	id := c.Query("ID")
	if err := auditAsRawService.DeleteAuditAsRaw(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAuditAsRawByIds 批量删除AS报文
// @Tags AuditAsRaw
// @Summary 批量删除AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /auditAsRaw/deleteAuditAsRawByIds [delete]
func (auditAsRawApi *AuditAsRawApi) DeleteAuditAsRawByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := auditAsRawService.DeleteAuditAsRawByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAuditAsRaw 更新AS报文
// @Tags AuditAsRaw
// @Summary 更新AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuditAsRaw true "更新AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /auditAsRaw/updateAuditAsRaw [put]
func (auditAsRawApi *AuditAsRawApi) UpdateAuditAsRaw(c *gin.Context) {
	var auditAsRaw model.AuditAsRaw
	err := c.ShouldBindJSON(&auditAsRaw)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := auditAsRawService.UpdateAuditAsRaw(auditAsRaw); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAuditAsRaw 用id查询AS报文
// @Tags AuditAsRaw
// @Summary 用id查询AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AuditAsRaw true "用id查询AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /auditAsRaw/findAuditAsRaw [get]
func (auditAsRawApi *AuditAsRawApi) FindAuditAsRaw(c *gin.Context) {
	id := c.Query("ID")
	if reauditAsRaw, err := auditAsRawService.GetAuditAsRaw(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reauditAsRaw": reauditAsRaw}, c)
	}
}

// GetAuditAsRawList 分页获取AS报文列表
// @Tags AuditAsRaw
// @Summary 分页获取AS报文列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AuditAsRawSearch true "分页获取AS报文列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /auditAsRaw/getAuditAsRawList [get]
func (auditAsRawApi *AuditAsRawApi) GetAuditAsRawList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AuditAsRawSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := auditAsRawService.GetAuditAsRawInfoList(pageInfo); err != nil {
		global.LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
