package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	request "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/service"
)

type AccountGsApi struct {
}

var accountGsService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AccountGsService

// CreateAccountGs 创建地面站
// @Tags AccountGs
// @Summary 创建地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountGs true "创建地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountGs/createAccountGs [post]
func (accountGsApi *AccountGsApi) CreateAccountGs(c *gin.Context) {
	var accountGs model.AccountGs
	err := c.ShouldBindJSON(&accountGs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountGsService.CreateAccountGs(&accountGs); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccountGs 删除地面站
// @Tags AccountGs
// @Summary 删除地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountGs true "删除地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountGs/deleteAccountGs [delete]
func (accountGsApi *AccountGsApi) DeleteAccountGs(c *gin.Context) {
	id := c.Query("ID")
	if err := accountGsService.DeleteAccountGs(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountGsByIds 批量删除地面站
// @Tags AccountGs
// @Summary 批量删除地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountGs/deleteAccountGsByIds [delete]
func (accountGsApi *AccountGsApi) DeleteAccountGsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := accountGsService.DeleteAccountGsByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccountGs 更新地面站
// @Tags AccountGs
// @Summary 更新地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountGs true "更新地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountGs/updateAccountGs [put]
func (accountGsApi *AccountGsApi) UpdateAccountGs(c *gin.Context) {
	var accountGs model.AccountGs
	err := c.ShouldBindJSON(&accountGs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountGsService.UpdateAccountGs(accountGs); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccountGs 用id查询地面站
// @Tags AccountGs
// @Summary 用id查询地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AccountGs true "用id查询地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountGs/findAccountGs [get]
func (accountGsApi *AccountGsApi) FindAccountGs(c *gin.Context) {
	id := c.Query("ID")
	if reaccountGs, err := accountGsService.GetAccountGs(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountGs": reaccountGs}, c)
	}
}

// GetAccountGsList 分页获取地面站列表
// @Tags AccountGs
// @Summary 分页获取地面站列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AccountGsSearch true "分页获取地面站列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountGs/getAccountGsList [get]
func (accountGsApi *AccountGsApi) GetAccountGsList(c *gin.Context) {
	var pageInfo request.AccountGsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountGsService.GetAccountGsInfoList(pageInfo); err != nil {
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
