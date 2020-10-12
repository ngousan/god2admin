package sysUserRoles

import "github.com/gin-gonic/gin"

func RolesRouter(Router *gin.RouterGroup) {
	routerRoles := Router.Group("roles")
	{
		routerRoles.GET("list", UserRolesList)
		routerRoles.POST("new", UserRolesAssign)
		routerRoles.POST("update", UserRolesUpdate)
	}
}
