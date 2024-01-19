package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AccountAsRouter struct {
}

// InitAccountAsRouter 初始化 飞机站账户 路由信息
func (s *AccountAsRouter) InitAccountAsRouter(Router *gin.RouterGroup) {
	accountAsRouter := Router.Group("accountAs").Use(middleware.OperationRecord())
	accountAsRouterWithoutRecord := Router.Group("accountAs")
	var accountAsApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AccountAsApi
	{
		accountAsRouter.POST("createAccountAs", accountAsApi.CreateAccountAs)             // 新建飞机站账户
		accountAsRouter.DELETE("deleteAccountAs", accountAsApi.DeleteAccountAs)           // 删除飞机站账户
		accountAsRouter.DELETE("deleteAccountAsByIds", accountAsApi.DeleteAccountAsByIds) // 批量删除飞机站账户
		accountAsRouter.PUT("updateAccountAs", accountAsApi.UpdateAccountAs)              // 更新飞机站账户
		accountAsRouter.PUT("setStateChange", accountAsApi.SetStateChange)                //启用状态更新
	}
	{
		accountAsRouterWithoutRecord.GET("findAccountAs", accountAsApi.FindAccountAs)       // 根据ID获取飞机站账户
		accountAsRouterWithoutRecord.GET("getAccountAsList", accountAsApi.GetAccountAsList) // 获取飞机站账户列表
		accountAsRouterWithoutRecord.GET("getOptions", accountAsApi.GetOptions)             //获取选项
	}
}
