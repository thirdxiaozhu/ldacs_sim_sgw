package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/pkg/forward_module/global"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
	"ldacs_sim_sgw/pkg/forward_module/service"
)

type AccountAuthzApi struct {
}

var accountAuthzService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AccountAuthzService

// CreateAccountAuthz 创建业务权限
// @Tags AccountAuthz
// @Summary 创建业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountAuthz true "创建业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountAuthz/createAccountAuthz [post]
func (accountAuthzApi *AccountAuthzApi) CreateAccountAuthz(c *gin.Context) {
	var accountAuthz ldacs_sgw_forward.AccountAuthz
	err := c.ShouldBindJSON(&accountAuthz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountAuthzService.CreateAccountAuthz(&accountAuthz); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccountAuthz 删除业务权限
// @Tags AccountAuthz
// @Summary 删除业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountAuthz true "删除业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountAuthz/deleteAccountAuthz [delete]
func (accountAuthzApi *AccountAuthzApi) DeleteAccountAuthz(c *gin.Context) {
	id := c.Query("ID")
	if err := accountAuthzService.DeleteAccountAuthz(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountAuthzByIds 批量删除业务权限
// @Tags AccountAuthz
// @Summary 批量删除业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountAuthz/deleteAccountAuthzByIds [delete]
func (accountAuthzApi *AccountAuthzApi) DeleteAccountAuthzByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := accountAuthzService.DeleteAccountAuthzByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccountAuthz 更新业务权限
// @Tags AccountAuthz
// @Summary 更新业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountAuthz true "更新业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountAuthz/updateAccountAuthz [put]
func (accountAuthzApi *AccountAuthzApi) UpdateAccountAuthz(c *gin.Context) {
	var accountAuthz ldacs_sgw_forward.AccountAuthz
	err := c.ShouldBindJSON(&accountAuthz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountAuthzService.UpdateAccountAuthz(accountAuthz); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccountAuthz 用id查询业务权限
// @Tags AccountAuthz
// @Summary 用id查询业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AccountAuthz true "用id查询业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountAuthz/findAccountAuthz [get]
func (accountAuthzApi *AccountAuthzApi) FindAccountAuthz(c *gin.Context) {
	id := c.Query("ID")
	if reaccountAuthz, err := accountAuthzService.GetAccountAuthz(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountAuthz": reaccountAuthz}, c)
	}
}

// GetAccountAuthzList 分页获取业务权限列表
// @Tags AccountAuthz
// @Summary 分页获取业务权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AccountAuthzSearch true "分页获取业务权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountAuthz/getAccountAuthzList [get]
func (accountAuthzApi *AccountAuthzApi) GetAccountAuthzList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AccountAuthzSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountAuthzService.GetAccountAuthzInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
