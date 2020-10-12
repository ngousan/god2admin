package login

import (
	"github.com/gin-gonic/gin"
)

func LoginRouter(Router *gin.RouterGroup) {
	routerLogin := Router.Group("")
	{
		routerLogin.POST("login", Login)
		routerLogin.POST("captcha", Captcha)
	}
}
