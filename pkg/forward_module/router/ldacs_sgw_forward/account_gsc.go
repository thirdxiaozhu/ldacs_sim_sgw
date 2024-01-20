package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AccountGscRouter struct {
}

// InitAccountGscRouter 初始化 地面控制站 路由信息
func (s *AccountGscRouter) InitAccountGscRouter(Router *gin.RouterGroup) {
	accountGscRouter := Router.Group("accountGsc").Use(middleware.OperationRecord())
	accountGscRouterWithoutRecord := Router.Group("accountGsc")
	var accountGscApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AccountGscApi
	{
		accountGscRouter.POST("createAccountGsc", accountGscApi.CreateAccountGsc)             // 新建地面控制站
		accountGscRouter.DELETE("deleteAccountGsc", accountGscApi.DeleteAccountGsc)           // 删除地面控制站
		accountGscRouter.DELETE("deleteAccountGscByIds", accountGscApi.DeleteAccountGscByIds) // 批量删除地面控制站
		accountGscRouter.PUT("updateAccountGsc", accountGscApi.UpdateAccountGsc)              // 更新地面控制站
	}
	{
		accountGscRouterWithoutRecord.GET("findAccountGsc", accountGscApi.FindAccountGsc)       // 根据ID获取地面控制站
		accountGscRouterWithoutRecord.GET("getAccountGscList", accountGscApi.GetAccountGscList) // 获取地面控制站列表
	}
}
