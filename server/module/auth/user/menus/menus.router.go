package authUserMenus

import (
	"github.com/gin-gonic/gin"
)

func UserMenusRouter(Router *gin.RouterGroup) {
	routerMenus := Router.Group("")
	{
		routerMenus.GET("menus", UserMenuList)
	}
}
