package sysRole

import (
	"god2admin/middleware"
	menus "god2admin/module/sys/role/menus"

	"github.com/gin-gonic/gin"
)

func RoleRouter(Router *gin.RouterGroup) {
	Router.Use(middleware.JWTAuth())
	routerRole := Router.Group("role")
	{
		routerRole.GET("list", RoleList)
		routerRole.POST("new", RoleCreate)
		routerRole.POST("update", RoleUpdate)
		routerRole.GET("dict", RoleDict)
		menus.MenusRouter(routerRole)
	}
}
