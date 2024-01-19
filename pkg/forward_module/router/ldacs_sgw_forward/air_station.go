package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AirStationRouter struct {
}

// InitAirStationRouter 初始化 飞机站 路由信息
func (s *AirStationRouter) InitAirStationRouter(Router *gin.RouterGroup) {
	airStationRouter := Router.Group("airStation").Use(middleware.OperationRecord())
	airStationRouterWithoutRecord := Router.Group("airStation")
	var airStationApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AirStationApi
	{
		airStationRouter.POST("createAirStation", airStationApi.CreateAirStation)             // 新建飞机站
		airStationRouter.DELETE("deleteAirStation", airStationApi.DeleteAirStation)           // 删除飞机站
		airStationRouter.DELETE("deleteAirStationByIds", airStationApi.DeleteAirStationByIds) // 批量删除飞机站
		airStationRouter.PUT("updateAirStation", airStationApi.UpdateAirStation)              // 更新飞机站
	}
	{
		airStationRouterWithoutRecord.GET("findAirStation", airStationApi.FindAirStation)       // 根据ID获取飞机站
		airStationRouterWithoutRecord.GET("getAirStationList", airStationApi.GetAirStationList) // 获取飞机站列表
	}
}
