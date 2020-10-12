package core

import (
	// _ "gin-vue-admin/docs"
	"god2admin/api"
	"god2admin/global"
	"god2admin/middleware"
	"god2admin/module/auth"
	"god2admin/module/sys"

	// "gin-vue-admin/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func initRouters() *gin.Engine {
	var app = gin.Default()
	app.Use(middleware.Logger()) // 日志
	// app.Use(middleware.LoadTls())  // 打开就能https了
	app.Use(middleware.Cors()) // 跨域
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// global.LOGGER.Debug("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	appRouter := app.Group("")
	{
		appRouter.GET("echo", api.Echo) //测试
		auth.AuthRouter(appRouter)
		sys.SysRouter(appRouter)
	}
	global.LOGGER.Info("router register success")
	return app
}
