package core

import (
	"fmt"
	"god2admin/global"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

const splitLine = "=================================================================="
const defaultConfigFile = "config/config.yaml"

type server interface {
	ListenAndServe() error
}

func init() {
	fmt.Println(splitLine)
	fmt.Println("初始化Config...")
	initConfig()
	fmt.Println("初始化Logger...")
	initLogger()
	fmt.Println("初始化Database...")
	initDatabase()
	fmt.Println("初始化Tables...")
	registerTables()
	fmt.Println(splitLine)

}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func StartServer() {
	fmt.Println(splitLine)
	global.LOGGER.Info("Start Server...")
	global.LOGGER.Info("初始化Router...")
	Router := initRouters()
	Router.Static("/form-generator", "./resource/page")
	global.LOGGER.Info("初始化Server...")
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Infof("欢迎使用 God2Admin!默认运行地址:http://127.0.0.1%s\n", address)
	fmt.Println(splitLine)
	global.LOGGER.Error(s.ListenAndServe())
}

func StopServer() {
	fmt.Println(splitLine)

	// 程序结束前关闭数据库链接
	global.LOGGER.Info("Close Database...")
	db, _ := global.DB.DB()
	defer db.Close()

	global.LOGGER.Info("Stop Server...")
	fmt.Println(splitLine)
}
