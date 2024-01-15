package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	v1 "ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AuthzPlaneRouter struct {
}

// InitAuthzPlaneRouter 初始化 飞机业务授权 路由信息
func (s *AuthzPlaneRouter) InitAuthzPlaneRouter(Router *gin.RouterGroup) {
	authzPlaneRouter := Router.Group("authzPlane").Use(middleware.OperationRecord())
	authzPlaneRouterWithoutRecord := Router.Group("authzPlane")
	var authzPlaneApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AuthzPlaneApi
	{
		authzPlaneRouter.POST("createAuthzPlane", authzPlaneApi.CreateAuthzPlane)             // 新建飞机业务授权
		authzPlaneRouter.DELETE("deleteAuthzPlane", authzPlaneApi.DeleteAuthzPlane)           // 删除飞机业务授权
		authzPlaneRouter.DELETE("deleteAuthzPlaneByIds", authzPlaneApi.DeleteAuthzPlaneByIds) // 批量删除飞机业务授权
		authzPlaneRouter.PUT("updateAuthzPlane", authzPlaneApi.UpdateAuthzPlane)              // 更新飞机业务授权
		authzPlaneRouter.PUT("setStateChange", authzPlaneApi.SetStateChange)                  //启用状态更新
	}
	{
		authzPlaneRouterWithoutRecord.GET("getOptions", authzPlaneApi.GetOptions)               //获取选项
		authzPlaneRouterWithoutRecord.GET("findAuthzPlane", authzPlaneApi.FindAuthzPlane)       // 根据ID获取飞机业务授权
		authzPlaneRouterWithoutRecord.GET("getAuthzPlaneList", authzPlaneApi.GetAuthzPlaneList) // 获取飞机业务授权列表
	}
}
