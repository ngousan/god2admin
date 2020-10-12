package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OkWithMessage(message string, c *gin.Context) {
	Result(0, map[string]interface{}{}, message, c)
}

func Echo(c *gin.Context) {
	// claims, _ := c.Get("claims")

	OkWithMessage(fmt.Sprintf("Get api"), c)
}
