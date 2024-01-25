package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ldacs_sim_sgw/internal/global"

	"ldacs_sim_sgw/pkg/forward_module/model/common/response"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
	"ldacs_sim_sgw/pkg/forward_module/utils"
)

type JwtApi struct{}

// JsonInBlacklist
// @Tags      Jwt
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "jwt加入黑名单"
// @Router    /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		global.LOGGER.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	utils.ClearToken(c)
	response.OkWithMessage("jwt作废成功", c)
}
