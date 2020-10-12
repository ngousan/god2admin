package request

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageListInfo(c *gin.Context) (pageListInfo PageListInfo, err error) {
	pageListInfo.Size, err = strconv.Atoi(c.Query("size"))
	pageListInfo.Current, err = strconv.Atoi(c.Query("current"))
	if pageListInfo.Size <= 0 || pageListInfo.Current <= 0 {
		err = errors.New(fmt.Sprintf("参数错误！size:%d; current:%d", pageListInfo.Size, pageListInfo.Current))
	}
	return
}

func GetClaims(c *gin.Context) (claims *CustomClaims) {
	cc, _ := c.Get("claims")
	claims = cc.(*CustomClaims)
	return
}
