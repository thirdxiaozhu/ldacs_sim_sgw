package ldacs_sgw_forward

import (
	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/api/v1"
	"ldacs_sim_sgw/pkg/forward_module/middleware"
)

type AuditAsRawRouter struct {
}

// InitAuditAsRawRouter 初始化 AS报文 路由信息
func (s *AuditAsRawRouter) InitAuditAsRawRouter(Router *gin.RouterGroup) {
	auditAsRawRouter := Router.Group("auditAsRaw").Use(middleware.OperationRecord())
	auditAsRawRouterWithoutRecord := Router.Group("auditAsRaw")
	var auditAsRawApi = v1.ApiGroupApp.Ldacs_sgw_forwardApiGroup.AuditAsRawApi
	{
		auditAsRawRouter.POST("createAuditAsRaw", auditAsRawApi.CreateAuditAsRaw)             // 新建AS报文
		auditAsRawRouter.DELETE("deleteAuditAsRaw", auditAsRawApi.DeleteAuditAsRaw)           // 删除AS报文
		auditAsRawRouter.DELETE("deleteAuditAsRawByIds", auditAsRawApi.DeleteAuditAsRawByIds) // 批量删除AS报文
		auditAsRawRouter.PUT("updateAuditAsRaw", auditAsRawApi.UpdateAuditAsRaw)              // 更新AS报文
	}
	{
		auditAsRawRouterWithoutRecord.GET("findAuditAsRaw", auditAsRawApi.FindAuditAsRaw)       // 根据ID获取AS报文
		auditAsRawRouterWithoutRecord.GET("getAuditAsRawList", auditAsRawApi.GetAuditAsRawList) // 获取AS报文列表
	}
}
