package utils

import (
	"fmt"
	"ldacs_sim_sgw/pkg/forward_module/f_global"

	"ldacs_sim_sgw/pkg/forward_module/model/system"
	"strconv"
)

func RegisterApis(apis ...system.SysApi) {
	var count int64
	var apiPaths []string
	for i := range apis {
		apiPaths = append(apiPaths, apis[i].Path)
	}
	global.DB.Find(&[]system.SysApi{}, "path in (?)", apiPaths).Count(&count)
	if count > 0 {
		fmt.Println("插件已安装或存在同名路由")
		return
	}
	err := global.DB.Create(&apis).Error
	if err != nil {
		fmt.Println(err)
	}
}

func RegisterMenus(menus ...system.SysBaseMenu) {
	var count int64
	var menuNames []string
	parentMenu := menus[0]
	otherMenus := menus[1:]
	for i := range menus {
		menuNames = append(menuNames, menus[i].Name)
	}
	global.DB.Find(&[]system.SysBaseMenu{}, "name in (?)", menuNames).Count(&count)
	if count > 0 {
		fmt.Println("插件已安装或存在同名菜单")
		return
	}
	parentMenu.ParentId = "0"
	err := global.DB.Create(&parentMenu).Error
	if err != nil {
		fmt.Println(err)
	}
	for i := range otherMenus {
		pid := strconv.Itoa(int(parentMenu.ID))
		otherMenus[i].ParentId = pid
	}
	err = global.DB.Create(&otherMenus).Error
	if err != nil {
		fmt.Println(err)
	}
}
