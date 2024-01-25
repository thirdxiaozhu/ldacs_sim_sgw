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

type AuthcStateApi struct {
}

var authcStateService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AuthcStateService

// CreateAuthcState 创建认证状态
// @Tags AuthcState
// @Summary 创建认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuthcState true "创建认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authcState/createAuthcState [post]
func (authcStateApi *AuthcStateApi) CreateAuthcState(c *gin.Context) {
	var authcState model.AuthcState
	err := c.ShouldBindJSON(&authcState)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := authcStateService.CreateAuthcState(&authcState); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAuthcState 删除认证状态
// @Tags AuthcState
// @Summary 删除认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuthcState true "删除认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authcState/deleteAuthcState [delete]
func (authcStateApi *AuthcStateApi) DeleteAuthcState(c *gin.Context) {
	id := c.Query("ID")
	if err := authcStateService.DeleteAuthcState(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAuthcStateByIds 批量删除认证状态
// @Tags AuthcState
// @Summary 批量删除认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /authcState/deleteAuthcStateByIds [delete]
func (authcStateApi *AuthcStateApi) DeleteAuthcStateByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := authcStateService.DeleteAuthcStateByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAuthcState 更新认证状态
// @Tags AuthcState
// @Summary 更新认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuthcState true "更新认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authcState/updateAuthcState [put]
func (authcStateApi *AuthcStateApi) UpdateAuthcState(c *gin.Context) {
	var authcState model.AuthcState
	err := c.ShouldBindJSON(&authcState)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := authcStateService.UpdateAuthcState(authcState); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAuthcState 用id查询认证状态
// @Tags AuthcState
// @Summary 用id查询认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AuthcState true "用id查询认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /authcState/findAuthcState [get]
func (authcStateApi *AuthcStateApi) FindAuthcState(c *gin.Context) {
	id := c.Query("ID")
	if reauthcState, err := authcStateService.GetAuthcState(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reauthcState": reauthcState}, c)
	}
}

// GetAuthcStateList 分页获取认证状态列表
// @Tags AuthcState
// @Summary 分页获取认证状态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AuthcStateSearch true "分页获取认证状态列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authcState/getAuthcStateList [get]
func (authcStateApi *AuthcStateApi) GetAuthcStateList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AuthcStateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := authcStateService.GetAuthcStateInfoList(pageInfo); err != nil {
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
