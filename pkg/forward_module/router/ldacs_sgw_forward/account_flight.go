package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AccontFlightRouter struct {
}

// InitAccontFlightRouter 初始化 航班 路由信息
func (s *AccontFlightRouter) InitAccontFlightRouter(Router *gin.RouterGroup) {
	accountFlightRouter := Router.Group("accountFlight").Use(middleware.OperationRecord())
	accountFlightRouterWithoutRecord := Router.Group("accountFlight")
	var accountFlightApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AccontFlightApi
	{
		accountFlightRouter.POST("createAccontFlight", accountFlightApi.CreateAccontFlight)             // 新建航班
		accountFlightRouter.DELETE("deleteAccontFlight", accountFlightApi.DeleteAccontFlight)           // 删除航班
		accountFlightRouter.DELETE("deleteAccontFlightByIds", accountFlightApi.DeleteAccontFlightByIds) // 批量删除航班
		accountFlightRouter.PUT("updateAccontFlight", accountFlightApi.UpdateAccontFlight)              // 更新航班
	}
	{
		accountFlightRouterWithoutRecord.GET("findAccontFlight", accountFlightApi.FindAccontFlight)       // 根据ID获取航班
		accountFlightRouterWithoutRecord.GET("getAccontFlightList", accountFlightApi.GetAccontFlightList) // 获取航班列表
	}
}
