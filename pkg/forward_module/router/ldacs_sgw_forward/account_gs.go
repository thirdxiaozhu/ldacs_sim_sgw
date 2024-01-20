package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AccountGsRouter struct {
}

// InitAccountGsRouter 初始化 地面站 路由信息
func (s *AccountGsRouter) InitAccountGsRouter(Router *gin.RouterGroup) {
	accountGsRouter := Router.Group("accountGs").Use(middleware.OperationRecord())
	accountGsRouterWithoutRecord := Router.Group("accountGs")
	var accountGsApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AccountGsApi
	{
		accountGsRouter.POST("createAccountGs", accountGsApi.CreateAccountGs)             // 新建地面站
		accountGsRouter.DELETE("deleteAccountGs", accountGsApi.DeleteAccountGs)           // 删除地面站
		accountGsRouter.DELETE("deleteAccountGsByIds", accountGsApi.DeleteAccountGsByIds) // 批量删除地面站
		accountGsRouter.PUT("updateAccountGs", accountGsApi.UpdateAccountGs)              // 更新地面站
	}
	{
		accountGsRouterWithoutRecord.GET("findAccountGs", accountGsApi.FindAccountGs)       // 根据ID获取地面站
		accountGsRouterWithoutRecord.GET("getAccountGsList", accountGsApi.GetAccountGsList) // 获取地面站列表
	}
}
