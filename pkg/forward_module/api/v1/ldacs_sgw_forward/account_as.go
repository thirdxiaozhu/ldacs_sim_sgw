package ldacs_sgw_forward

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/service"
	"ldacs_sim_sgw/pkg/forward_module/utils"
)

type AccountAsApi struct {
}

var accountAsService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AccountAsService

// CreateAccountAs 创建飞机站账户
// @Tags AccountAs
// @Summary 创建飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountAs true "创建飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountAs/createAccountAs [post]
func (accountAsApi *AccountAsApi) CreateAccountAs(c *gin.Context) {
	var accountAs model.AccountAs
	err := c.ShouldBindJSON(&accountAs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	accountAs.CreatedBy = utils.GetUserID(c)

	if err := accountAsService.CreateAccountAs(&accountAs); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccountAs 删除飞机站账户
// @Tags AccountAs
// @Summary 删除飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountAs true "删除飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountAs/deleteAccountAs [delete]
func (accountAsApi *AccountAsApi) DeleteAccountAs(c *gin.Context) {
	id := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := accountAsService.DeleteAccountAs(id, userID); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountAsByIds 批量删除飞机站账户
// @Tags AccountAs
// @Summary 批量删除飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountAs/deleteAccountAsByIds [delete]
func (accountAsApi *AccountAsApi) DeleteAccountAsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := accountAsService.DeleteAccountAsByIds(ids, userID); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccountAs 更新飞机站账户
// @Tags AccountAs
// @Summary 更新飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccountAs true "更新飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountAs/updateAccountAs [put]
func (accountAsApi *AccountAsApi) UpdateAccountAs(c *gin.Context) {
	var accountAs model.AccountAs
	err := c.ShouldBindJSON(&accountAs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	accountAs.UpdatedBy = utils.GetUserID(c)

	if err := accountAsService.UpdateAccountAs(accountAs); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccountAs 用id查询飞机站账户
// @Tags AccountAs
// @Summary 用id查询飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AccountAs true "用id查询飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountAs/findAccountAs [get]
func (accountAsApi *AccountAsApi) FindAccountAs(c *gin.Context) {
	id := c.Query("ID")
	if reaccountAs, err := accountAsService.GetAccountAs(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountAs": reaccountAs}, c)
	}
}

// GetAccountAsList 分页获取飞机站账户列表
// @Tags AccountAs
// @Summary 分页获取飞机站账户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AccountAsSearch true "分页获取飞机站账户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountAs/getAccountAsList [get]
func (accountAsApi *AccountAsApi) GetAccountAsList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AccountAsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountAsService.GetAccountAsInfoList(pageInfo); err != nil {
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
func (accountAsApi *AccountAsApi) GetOptions(c *gin.Context) {
	//err := c.ShouldBindQuery(&troq)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}

	if opts, err := accountAsService.GetOptions(); err != nil {
		global.LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"options": opts}, c)
	}
}

func (accountAsApi *AccountAsApi) SetStateChange(c *gin.Context) {
	var accountAs model.AccountAs
	err := c.ShouldBindJSON(&accountAs)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Printf("STATE %d\n", accountAs.AsCurrState)

	if err := accountAsService.StateChange(&accountAs); err != nil {
		global.LOGGER.Error("授权启动失败!", zap.Error(err))
		response.FailWithMessage("授权启动失败", c)
	} else {
		retMap := gin.H{}
		retMap["state"] = accountAs.AsCurrState
		//if accountAs.AsState == 1 {
		//} else {
		//	retMap["state"] = 0
		//}
		response.OkWithData(retMap, c)
	}
}
