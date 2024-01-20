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

type AccountGscApi struct {
}

var accountGscService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AccountGscService

// CreateAccountGsc 创建地面控制站
// @Tags AccountGsc
// @Summary 创建地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountGsc true "创建地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountGsc/createAccountGsc [post]
func (accountGscApi *AccountGscApi) CreateAccountGsc(c *gin.Context) {
	var accountGsc model.AccountGsc
	err := c.ShouldBindJSON(&accountGsc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountGscService.CreateAccountGsc(&accountGsc); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccountGsc 删除地面控制站
// @Tags AccountGsc
// @Summary 删除地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountGsc true "删除地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountGsc/deleteAccountGsc [delete]
func (accountGscApi *AccountGscApi) DeleteAccountGsc(c *gin.Context) {
	id := c.Query("ID")
	if err := accountGscService.DeleteAccountGsc(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountGscByIds 批量删除地面控制站
// @Tags AccountGsc
// @Summary 批量删除地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountGsc/deleteAccountGscByIds [delete]
func (accountGscApi *AccountGscApi) DeleteAccountGscByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := accountGscService.DeleteAccountGscByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccountGsc 更新地面控制站
// @Tags AccountGsc
// @Summary 更新地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountGsc true "更新地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountGsc/updateAccountGsc [put]
func (accountGscApi *AccountGscApi) UpdateAccountGsc(c *gin.Context) {
	var accountGsc model.AccountGsc
	err := c.ShouldBindJSON(&accountGsc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountGscService.UpdateAccountGsc(accountGsc); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccountGsc 用id查询地面控制站
// @Tags AccountGsc
// @Summary 用id查询地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AccountGsc true "用id查询地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountGsc/findAccountGsc [get]
func (accountGscApi *AccountGscApi) FindAccountGsc(c *gin.Context) {
	id := c.Query("ID")
	if reaccountGsc, err := accountGscService.GetAccountGsc(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountGsc": reaccountGsc}, c)
	}
}

// GetAccountGscList 分页获取地面控制站列表
// @Tags AccountGsc
// @Summary 分页获取地面控制站列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AccountGscSearch true "分页获取地面控制站列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountGsc/getAccountGscList [get]
func (accountGscApi *AccountGscApi) GetAccountGscList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AccountGscSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountGscService.GetAccountGscInfoList(pageInfo); err != nil {
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
