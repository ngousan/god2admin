package authUser

import (
	"god2admin/middleware"
	menus "god2admin/module/auth/user/menus"
	routers "god2admin/module/auth/user/routers"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	Router.Use(middleware.JWTAuth())
	routerUser := Router.Group("user")
	{
		menus.UserMenusRouter(routerUser)
		routers.UserRoutersRouter(routerUser)
	}
}
