package sysUser

import (
	"encoding/json"
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	page, err := request.GetPageListInfo(c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	var sysUser SysUser
	sysUser.Username = c.Query("username")
	userList, total, err := userList(page, sysUser)
	_, err = json.Marshal(userList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedPageList(page, total, userList, c)
}

func UserRegister(c *gin.Context) {
	var sysUser SysUser
	_ = c.ShouldBindJSON(&sysUser)
	err := userRegister(sysUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func UserUpdate(c *gin.Context) {
	var sysUser SysUser
	_ = c.ShouldBindJSON(&sysUser)
	err := userUpdate(&sysUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func UserDict(c *gin.Context) {
	ud, err := userDict()
	_, err = json.Marshal(ud)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedDict(ud, c)
}
