package ldacs_sgw_forward

import (
	"fmt"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
	ldacs_sgw_forwardReq "ldacs_sim_sgw/pkg/ldacs_core/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"

	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/service"
)

type AuthzPlaneApi struct {
}

var authzPlaneService = service.ServiceGroupApp.Ldacs_sgw_forwardServiceGroup.AuthzPlaneService

// CreateAuthzPlane 创建飞机业务授权
// @Tags AuthzPlane
// @Summary 创建飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuthzPlane true "创建飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authzPlane/createAuthzPlane [post]
func (authzPlaneApi *AuthzPlaneApi) CreateAuthzPlane(c *gin.Context) {
	//var authzPlane ldacs_sgw_forward.AuthzPlane
	var authzPlaneMulti model.AuthzPlaneMulti
	err := c.ShouldBindJSON(&authzPlaneMulti)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(authzPlaneMulti)

	for _, authz_n := range authzPlaneMulti.AuthzAuthzs {
		authzPlaneSingal := model.AuthzPlane{
			AuthzAs:    authzPlaneMulti.AuthzAs,
			AuthzAuthz: authz_n,
		}

		if err := authzPlaneService.CreateAuthzPlane(&authzPlaneSingal); err != nil {
			global.LOGGER.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		}
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAuthzPlane 删除飞机业务授权
// @Tags AuthzPlane
// @Summary 删除飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuthzPlane true "删除飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authzPlane/deleteAuthzPlane [delete]
func (authzPlaneApi *AuthzPlaneApi) DeleteAuthzPlane(c *gin.Context) {
	id := c.Query("ID")
	if err := authzPlaneService.DeleteAuthzPlane(id); err != nil {
		global.LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAuthzPlaneByIds 批量删除飞机业务授权
// @Tags AuthzPlane
// @Summary 批量删除飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /authzPlane/deleteAuthzPlaneByIds [delete]
func (authzPlaneApi *AuthzPlaneApi) DeleteAuthzPlaneByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := authzPlaneService.DeleteAuthzPlaneByIds(ids); err != nil {
		global.LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAuthzPlane 更新飞机业务授权
// @Tags AuthzPlane
// @Summary 更新飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ldacs_sgw_forward.AuthzPlane true "更新飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authzPlane/updateAuthzPlane [put]
func (authzPlaneApi *AuthzPlaneApi) UpdateAuthzPlane(c *gin.Context) {
	var authzPlane model.AuthzPlane
	err := c.ShouldBindJSON(&authzPlane)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := authzPlaneService.UpdateAuthzPlane(authzPlane); err != nil {
		global.LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAuthzPlane 用id查询飞机业务授权
// @Tags AuthzPlane
// @Summary 用id查询飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forward.AuthzPlane true "用id查询飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /authzPlane/findAuthzPlane [get]
func (authzPlaneApi *AuthzPlaneApi) FindAuthzPlane(c *gin.Context) {
	id := c.Query("ID")
	if reauthzPlane, err := authzPlaneService.GetAuthzPlane(id); err != nil {
		global.LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reauthzPlane": reauthzPlane}, c)
	}
}

// GetAuthzPlaneList 分页获取飞机业务授权列表
// @Tags AuthzPlane
// @Summary 分页获取飞机业务授权列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ldacs_sgw_forwardReq.AuthzPlaneSearch true "分页获取飞机业务授权列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authzPlane/getAuthzPlaneList [get]
func (authzPlaneApi *AuthzPlaneApi) GetAuthzPlaneList(c *gin.Context) {
	var pageInfo ldacs_sgw_forwardReq.AuthzPlaneSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := authzPlaneService.GetAuthzPlaneInfoList(pageInfo); err != nil {
		global.LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (authzPlaneApi *AuthzPlaneApi) GetOptions(c *gin.Context) {
	//err := c.ShouldBindQuery(&troq)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}

	if opts, err := authzPlaneService.GetOptions(); err != nil {
		global.LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"options": opts}, c)
	}
}

func (authzPlaneApi *AuthzPlaneApi) SetStateChange(c *gin.Context) {
	var authzPlane model.AuthzPlane
	err := c.ShouldBindJSON(&authzPlane)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Printf("STATE %d\n", authzPlane.AuthzState)

	if err := authzPlaneService.StateChange(&authzPlane); err != nil {
		global.LOGGER.Error("授权启动失败!", zap.Error(err))
		response.FailWithMessage("授权启动失败", c)
	} else {
		retMap := gin.H{}
		if authzPlane.AuthzState == 1 {
			retMap["state"] = 1
		} else {
			retMap["state"] = 0
		}
		response.OkWithData(retMap, c)
	}
}
