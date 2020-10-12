package sysMenu

import (
	"github.com/gin-gonic/gin"
)

func MenuRouter(Router *gin.RouterGroup) {
	routerMenu := Router.Group("menu")
	{
		routerMenu.GET("list", List)
		routerMenu.POST("new", Create)
		routerMenu.POST("update", Update)
		routerMenu.GET("dict", MenuDict)
	}
}
