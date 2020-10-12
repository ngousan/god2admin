package response

import (
	// modelResponse "god2admin/model/response"
	"god2admin/public/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func SucceedPageList(page request.PageListInfo, total int64, records interface{}, c *gin.Context) {
	Result(SUCCESS, ResponseData{
		Data: ListRecords{
			Current: page.Current,
			Size:    page.Size,
			Total:   total,
			Records: records,
		},
	}, "获取页面数据成功", c)
}

func SucceedPageListTree(page request.PageListInfo, total int64, records interface{}, c *gin.Context) {
	Result(SUCCESS, ResponseData{
		Data: ListTreeRecords{
			Records: records,
			Total:   total,
			Size:    page.Size,
			Current: page.Current,
		},
	}, "获取页面数据成功", c)
}

func SucceedDict(dict interface{}, c *gin.Context) {
	Result(SUCCESS, ResponseData{
		Data: dict,
	}, "获取字典成功", c)
}
