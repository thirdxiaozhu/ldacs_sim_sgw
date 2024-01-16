package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
	"ldacs_sim_sgw/pkg/forward_module/service"
)

type AccountPlaneApi struct {
}

var accountplaneService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AccountPlaneService

// CreateAccountPlane 创建飞机账户管理
// @Tags AccountPlane
// @Summary 创建飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountPlane true "创建飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountplane/createAccountPlane [post]
func (accountplaneApi *AccountPlaneApi) CreateAccountPlane(c *gin.Context) {
	var accountplane ldacs_sgw_forward.AccountPlane
	err := c.ShouldBindJSON(&accountplane)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountplaneService.CreateAccountPlane(&accountplane); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccountPlane 删除飞机账户管理
// @Tags AccountPlane
// @Summary 删除飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountPlane true "删除飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountplane/deleteAccountPlane [delete]
func (accountplaneApi *AccountPlaneApi) DeleteAccountPlane(c *gin.Context) {
	id := c.Query("ID")
	if err := accountplaneService.DeleteAccountPlane(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountPlaneByIds 批量删除飞机账户管理
// @Tags AccountPlane
// @Summary 批量删除飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountplane/deleteAccountPlaneByIds [delete]
func (accountplaneApi *AccountPlaneApi) DeleteAccountPlaneByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := accountplaneService.DeleteAccountPlaneByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccountPlane 更新飞机账户管理
// @Tags AccountPlane
// @Summary 更新飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountPlane true "更新飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountplane/updateAccountPlane [put]
func (accountplaneApi *AccountPlaneApi) UpdateAccountPlane(c *gin.Context) {
	var accountplane ldacs_sgw_forward.AccountPlane
	err := c.ShouldBindJSON(&accountplane)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountplaneService.UpdateAccountPlane(accountplane); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccountPlane 用id查询飞机账户管理
// @Tags AccountPlane
// @Summary 用id查询飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AccountPlane true "用id查询飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountplane/findAccountPlane [get]
func (accountplaneApi *AccountPlaneApi) FindAccountPlane(c *gin.Context) {
	id := c.Query("ID")
	if reaccountplane, err := accountplaneService.GetAccountPlane(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountplane": reaccountplane}, c)
	}
}

// GetAccountPlaneList 分页获取飞机账户管理列表
// @Tags AccountPlane
// @Summary 分页获取飞机账户管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AccountPlaneSearch true "分页获取飞机账户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountplane/getAccountPlaneList [get]
func (accountplaneApi *AccountPlaneApi) GetAccountPlaneList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AccountPlaneSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountplaneService.GetAccountPlaneInfoList(pageInfo); err != nil {
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
