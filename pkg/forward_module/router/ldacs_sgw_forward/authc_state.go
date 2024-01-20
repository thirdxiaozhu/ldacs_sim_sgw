package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AuthcStateRouter struct {
}

// InitAuthcStateRouter 初始化 认证状态 路由信息
func (s *AuthcStateRouter) InitAuthcStateRouter(Router *gin.RouterGroup) {
	authcStateRouter := Router.Group("authcState").Use(middleware.OperationRecord())
	authcStateRouterWithoutRecord := Router.Group("authcState")
	var authcStateApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AuthcStateApi
	{
		authcStateRouter.POST("createAuthcState", authcStateApi.CreateAuthcState)             // 新建认证状态
		authcStateRouter.DELETE("deleteAuthcState", authcStateApi.DeleteAuthcState)           // 删除认证状态
		authcStateRouter.DELETE("deleteAuthcStateByIds", authcStateApi.DeleteAuthcStateByIds) // 批量删除认证状态
		authcStateRouter.PUT("updateAuthcState", authcStateApi.UpdateAuthcState)              // 更新认证状态
	}
	{
		authcStateRouterWithoutRecord.GET("findAuthcState", authcStateApi.FindAuthcState)       // 根据ID获取认证状态
		authcStateRouterWithoutRecord.GET("getAuthcStateList", authcStateApi.GetAuthcStateList) // 获取认证状态列表
	}
}
