package sys

import (
	menu "god2admin/module/sys/menu"
	role "god2admin/module/sys/role"
	user "god2admin/module/sys/user"

	"github.com/gin-gonic/gin"
)

func SysRouter(Router *gin.RouterGroup) {
	routerSys := Router.Group("sys")
	{
		user.UserRouter(routerSys)
		menu.MenuRouter(routerSys)
		role.RoleRouter(routerSys)
	}
}
