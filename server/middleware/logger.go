package middleware

import (
	"fmt"
	"god2admin/global"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {

	logger := global.LOGGER

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	//设置输出
	// logger.Out = loggerToFile()

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		logger.Infof("| %19s | %3d | %15s | %6s | %s | %8v |",
			startTime.Format("2006/01/02 15:04:05"),
			c.Writer.Status(),    // 状态码
			c.ClientIP(),         // 请求IP
			c.Request.Method,     // 请求方式
			c.Request.RequestURI, // 请求路由
			latencyTime,
		)
	}
}

// 日志记录到文件
func loggerToFile() *os.File {

	logFilePath := "/"
	logFileName := "god2admin.log"

	//日志文件
	logFile := path.Join(logFilePath, logFileName)

	//写入文件
	log, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	return log

}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
