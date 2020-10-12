package auth

import (
	login "god2admin/module/auth/login"
	user "god2admin/module/auth/user"

	"github.com/gin-gonic/gin"
)

func AuthRouter(Router *gin.RouterGroup) {
	routerAuth := Router.Group("auth")
	{
		login.LoginRouter(routerAuth)
		user.UserRouter(routerAuth)
	}
}
