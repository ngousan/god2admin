package sysRoleMenus

import (
	"github.com/gin-gonic/gin"
)

func MenusRouter(Router *gin.RouterGroup) {
	routerMenus := Router.Group("menus")
	{
		routerMenus.GET("list", RoleMenusList)
		routerMenus.POST("new", RoleMenusAssign)
		routerMenus.POST("update", RoleMenusUpdate)
	}
}
