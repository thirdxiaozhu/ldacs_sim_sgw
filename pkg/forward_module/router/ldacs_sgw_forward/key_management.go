package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type KeyEntityRouter struct {
}

// InitKeyEntityRouter 初始化 密钥 路由信息
func (s *KeyEntityRouter) InitKeyEntityRouter(Router *gin.RouterGroup) {
	kmRouter := Router.Group("km").Use(middleware.OperationRecord())
	kmRouterWithoutRecord := Router.Group("km")
	var kmApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.KeyEntityApi
	{
		kmRouter.POST("createKeyEntity", kmApi.CreateKeyEntity)             // 新建密钥
		kmRouter.DELETE("deleteKeyEntity", kmApi.DeleteKeyEntity)           // 删除密钥
		kmRouter.DELETE("deleteKeyEntityByIds", kmApi.DeleteKeyEntityByIds) // 批量删除密钥
		kmRouter.PUT("updateKeyEntity", kmApi.UpdateKeyEntity)              // 更新密钥
	}
	{
		kmRouterWithoutRecord.GET("findKeyEntity", kmApi.FindKeyEntity)       // 根据ID获取密钥
		kmRouterWithoutRecord.GET("getKeyEntityList", kmApi.GetKeyEntityList) // 获取密钥列表
		kmRouterWithoutRecord.GET("getOptions", kmApi.GetOptions)             // 获取密钥列表
	}
}
