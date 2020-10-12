package sysRole

import (
	"encoding/json"
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"

	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	page, err := request.GetPageListInfo(c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	roleList, total, err := list(page)
	_, err = json.Marshal(roleList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedPageList(page, total, roleList, c)
}

func RoleCreate(c *gin.Context) {
	var sysRole SysRole
	_ = c.ShouldBindJSON(&sysRole)
	err := insert(sysRole)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func RoleUpdate(c *gin.Context) {
	var sysRole SysRole
	_ = c.ShouldBindJSON(&sysRole)
	err := update(&sysRole)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func RoleDict(c *gin.Context) {
	rd, err := roleDict()
	_, err = json.Marshal(rd)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedDict(rd, c)
}
