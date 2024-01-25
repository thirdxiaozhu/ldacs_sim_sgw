package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AccountAuthzRouter struct {
}

// InitAccountAuthzRouter 初始化 业务权限 路由信息
func (s *AccountAuthzRouter) InitAccountAuthzRouter(Router *gin.RouterGroup) {
	accountAuthzRouter := Router.Group("accountAuthz").Use(middleware.OperationRecord())
	accountAuthzRouterWithoutRecord := Router.Group("accountAuthz")
	var accountAuthzApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AccountAuthzApi
	{
		accountAuthzRouter.POST("createAccountAuthz", accountAuthzApi.CreateAccountAuthz)             // 新建业务权限
		accountAuthzRouter.DELETE("deleteAccountAuthz", accountAuthzApi.DeleteAccountAuthz)           // 删除业务权限
		accountAuthzRouter.DELETE("deleteAccountAuthzByIds", accountAuthzApi.DeleteAccountAuthzByIds) // 批量删除业务权限
		accountAuthzRouter.PUT("updateAccountAuthz", accountAuthzApi.UpdateAccountAuthz)              // 更新业务权限
	}
	{
		accountAuthzRouterWithoutRecord.GET("findAccountAuthz", accountAuthzApi.FindAccountAuthz)       // 根据ID获取业务权限
		accountAuthzRouterWithoutRecord.GET("getAccountAuthzList", accountAuthzApi.GetAccountAuthzList) // 获取业务权限列表
	}
}
