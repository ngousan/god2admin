package authUserRouters

import (
	"github.com/gin-gonic/gin"
)

func UserRoutersRouter(Router *gin.RouterGroup) {
	routerRouters := Router.Group("")
	{
		routerRouters.GET("routers", UserRouterList)
	}
}
