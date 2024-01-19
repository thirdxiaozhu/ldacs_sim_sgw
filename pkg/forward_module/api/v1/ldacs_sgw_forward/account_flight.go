package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/service"
)

type AccontFlightApi struct {
}

var accountFlightService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AccontFlightService

// CreateAccontFlight 创建航班
// @Tags AccontFlight
// @Summary 创建航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccontFlight true "创建航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountFlight/createAccontFlight [post]
func (accountFlightApi *AccontFlightApi) CreateAccontFlight(c *gin.Context) {
	var accountFlight model.AccountFlight
	err := c.ShouldBindJSON(&accountFlight)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountFlightService.CreateAccontFlight(&accountFlight); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccontFlight 删除航班
// @Tags AccontFlight
// @Summary 删除航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccontFlight true "删除航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountFlight/deleteAccontFlight [delete]
func (accountFlightApi *AccontFlightApi) DeleteAccontFlight(c *gin.Context) {
	id := c.Query("ID")
	if err := accountFlightService.DeleteAccontFlight(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccontFlightByIds 批量删除航班
// @Tags AccontFlight
// @Summary 批量删除航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountFlight/deleteAccontFlightByIds [delete]
func (accountFlightApi *AccontFlightApi) DeleteAccontFlightByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := accountFlightService.DeleteAccontFlightByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccontFlight 更新航班
// @Tags AccontFlight
// @Summary 更新航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AccontFlight true "更新航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountFlight/updateAccontFlight [put]
func (accountFlightApi *AccontFlightApi) UpdateAccontFlight(c *gin.Context) {
	var accountFlight model.AccountFlight
	err := c.ShouldBindJSON(&accountFlight)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := accountFlightService.UpdateAccontFlight(accountFlight); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccontFlight 用id查询航班
// @Tags AccontFlight
// @Summary 用id查询航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AccontFlight true "用id查询航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountFlight/findAccontFlight [get]
func (accountFlightApi *AccontFlightApi) FindAccontFlight(c *gin.Context) {
	id := c.Query("ID")
	if reaccountFlight, err := accountFlightService.GetAccontFlight(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountFlight": reaccountFlight}, c)
	}
}

// GetAccontFlightList 分页获取航班列表
// @Tags AccontFlight
// @Summary 分页获取航班列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AccontFlightSearch true "分页获取航班列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountFlight/getAccontFlightList [get]
func (accountFlightApi *AccontFlightApi) GetAccontFlightList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AccontFlightSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountFlightService.GetAccontFlightInfoList(pageInfo); err != nil {
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
