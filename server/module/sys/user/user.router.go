package sysUser

import (
	roles "god2admin/module/sys/user/roles"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	routerUser := Router.Group("user")
	{
		routerUser.GET("list", UserList)
		routerUser.POST("new", UserRegister)
		routerUser.POST("update", UserUpdate)
		routerUser.GET("dict", UserDict)
		roles.RolesRouter(routerUser)
	}
}
