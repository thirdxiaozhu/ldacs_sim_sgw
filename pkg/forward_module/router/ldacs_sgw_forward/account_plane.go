package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AccountPlaneRouter struct {
}

// InitAccountPlaneRouter 初始化 飞机账户管理 路由信息
func (s *AccountPlaneRouter) InitAccountPlaneRouter(Router *gin.RouterGroup) {
	accountplaneRouter := Router.Group("accountplane").Use(middleware.OperationRecord())
	accountplaneRouterWithoutRecord := Router.Group("accountplane")
	var accountplaneApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AccountPlaneApi
	{
		accountplaneRouter.POST("createAccountPlane", accountplaneApi.CreateAccountPlane)             // 新建飞机账户管理
		accountplaneRouter.DELETE("deleteAccountPlane", accountplaneApi.DeleteAccountPlane)           // 删除飞机账户管理
		accountplaneRouter.DELETE("deleteAccountPlaneByIds", accountplaneApi.DeleteAccountPlaneByIds) // 批量删除飞机账户管理
		accountplaneRouter.PUT("updateAccountPlane", accountplaneApi.UpdateAccountPlane)              // 更新飞机账户管理
	}
	{
		accountplaneRouterWithoutRecord.GET("findAccountPlane", accountplaneApi.FindAccountPlane)       // 根据ID获取飞机账户管理
		accountplaneRouterWithoutRecord.GET("getAccountPlaneList", accountplaneApi.GetAccountPlaneList) // 获取飞机账户管理列表
	}
}
