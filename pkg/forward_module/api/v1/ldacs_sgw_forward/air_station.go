package ldacs_sgw_forward

import (
	"ldacs_sim_sgw/internal/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/forward_module/model/ldacs_sgw_forward/request"
	"ldacs_sim_sgw/pkg/forward_module/service"
	"ldacs_sim_sgw/pkg/forward_module/utils"
)

type AirStationApi struct {
}

var airStationService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AirStationService

// CreateAirStation 创建飞机站
// @Tags AirStation
// @Summary 创建飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AirStation true "创建飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /airStation/createAirStation [post]
func (airStationApi *AirStationApi) CreateAirStation(c *gin.Context) {
	var airStation ldacs_sgw_forward.AirStation
	err := c.ShouldBindJSON(&airStation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	airStation.CreatedBy = utils.GetUserID(c)

	if err := airStationService.CreateAirStation(&airStation); err != nil {
		global.LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAirStation 删除飞机站
// @Tags AirStation
// @Summary 删除飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AirStation true "删除飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /airStation/deleteAirStation [delete]
func (airStationApi *AirStationApi) DeleteAirStation(c *gin.Context) {
	id := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := airStationService.DeleteAirStation(id, userID); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAirStationByIds 批量删除飞机站
// @Tags AirStation
// @Summary 批量删除飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /airStation/deleteAirStationByIds [delete]
func (airStationApi *AirStationApi) DeleteAirStationByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := airStationService.DeleteAirStationByIds(ids, userID); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAirStation 更新飞机站
// @Tags AirStation
// @Summary 更新飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AirStation true "更新飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /airStation/updateAirStation [put]
func (airStationApi *AirStationApi) UpdateAirStation(c *gin.Context) {
	var airStation ldacs_sgw_forward.AirStation
	err := c.ShouldBindJSON(&airStation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	airStation.UpdatedBy = utils.GetUserID(c)

	if err := airStationService.UpdateAirStation(airStation); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAirStation 用id查询飞机站
// @Tags AirStation
// @Summary 用id查询飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AirStation true "用id查询飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /airStation/findAirStation [get]
func (airStationApi *AirStationApi) FindAirStation(c *gin.Context) {
	id := c.Query("ID")
	if reairStation, err := airStationService.GetAirStation(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reairStation": reairStation}, c)
	}
}

// GetAirStationList 分页获取飞机站列表
// @Tags AirStation
// @Summary 分页获取飞机站列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AirStationSearch true "分页获取飞机站列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /airStation/getAirStationList [get]
func (airStationApi *AirStationApi) GetAirStationList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AirStationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := airStationService.GetAirStationInfoList(pageInfo); err != nil {
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
