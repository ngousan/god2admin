package sysUserRoles

import (
	"encoding/json"
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 角色分配

func UserRolesList(c *gin.Context) {
	page, err := request.GetPageListInfo(c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	searchParam := make(map[string]interface{})
	searchUserId, err := strconv.Atoi(c.Query("userId"))
	if searchUserId != 0 {
		searchParam["user_id"] = searchUserId
	}
	// 	searchParam["user_id"] = "%"
	// }
	userList, total, err := userRolesList(page, searchParam)
	_, err = json.Marshal(userList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedPageList(page, total, userList, c)
}

func UserRolesAssign(c *gin.Context) {
	var sura SysUserRolesAssign
	_ = c.ShouldBindJSON(&sura)
	err := userRolesAssign(sura)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func UserRolesUpdate(c *gin.Context) {
	var sysUserRoles SysUserRoles
	_ = c.ShouldBindJSON(&sysUserRoles)
	fmt.Println(sysUserRoles)
	err := userRolesUpdate(&sysUserRoles)
	// var err error
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
