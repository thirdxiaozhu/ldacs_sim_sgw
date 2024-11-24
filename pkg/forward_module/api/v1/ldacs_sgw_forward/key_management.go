package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/service"
	"ldacs_sim_sgw/pkg/forward_module/utils"
)

type KeyEntityApi struct {
}

var kmService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.KeyEntityService

// CreateKeyEntity 创建密钥
// @Tags KeyEntity
// @Summary 创建密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.KeyEntity true "创建密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /km/createKeyEntity [post]
func (kmApi *KeyEntityApi) CreateKeyEntity(c *gin.Context) {
	var km model.KeyEntity
	err := c.ShouldBindJSON(&km)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	km.CreatedBy = utils.GetUserID(c)

	if err := kmService.CreateKeyEntity(&km); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKeyEntity 删除密钥
// @Tags KeyEntity
// @Summary 删除密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.KeyEntity true "删除密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /km/deleteKeyEntity [delete]
func (kmApi *KeyEntityApi) DeleteKeyEntity(c *gin.Context) {
	id := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := kmService.DeleteKeyEntity(id, userID); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKeyEntityByIds 批量删除密钥
// @Tags KeyEntity
// @Summary 批量删除密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /km/deleteKeyEntityByIds [delete]
func (kmApi *KeyEntityApi) DeleteKeyEntityByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := kmService.DeleteKeyEntityByIds(ids, userID); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKeyEntity 更新密钥
// @Tags KeyEntity
// @Summary 更新密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.KeyEntity true "更新密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /km/updateKeyEntity [put]
func (kmApi *KeyEntityApi) UpdateKeyEntity(c *gin.Context) {
	var km model.KeyEntity
	err := c.ShouldBindJSON(&km)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	km.UpdatedBy = utils.GetUserID(c)

	if err := kmService.UpdateKeyEntity(km); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKeyEntity 用id查询密钥
// @Tags KeyEntity
// @Summary 用id查询密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.KeyEntity true "用id查询密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /km/findKeyEntity [get]
func (kmApi *KeyEntityApi) FindKeyEntity(c *gin.Context) {
	id := c.Query("ID")
	if rekm, err := kmService.GetKeyEntity(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rekm": rekm}, c)
	}
}

// GetKeyEntityList 分页获取密钥列表
// @Tags KeyEntity
// @Summary 分页获取密钥列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.KeyEntitySearch true "分页获取密钥列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /km/getKeyEntityList [get]
func (kmApi *KeyEntityApi) GetKeyEntityList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.KeyEntitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := kmService.GetKeyEntityInfoList(pageInfo); err != nil {
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
