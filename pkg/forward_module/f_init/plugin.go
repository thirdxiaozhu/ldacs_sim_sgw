package f_init

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"ldacs_sim_sgw/pkg/forward_module/f_global"

	"ldacs_sim_sgw/pkg/forward_module/middleware"
	"ldacs_sim_sgw/pkg/forward_module/plugin/email"
	"ldacs_sim_sgw/pkg/forward_module/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		f_global.GVA_CONFIG.Email.To,
		f_global.GVA_CONFIG.Email.From,
		f_global.GVA_CONFIG.Email.Host,
		f_global.GVA_CONFIG.Email.Secret,
		f_global.GVA_CONFIG.Email.Nickname,
		f_global.GVA_CONFIG.Email.Port,
		f_global.GVA_CONFIG.Email.IsSSL,
	))
}
