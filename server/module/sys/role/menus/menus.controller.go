package sysRoleMenus

import (
	"encoding/json"
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 菜单分配

func RoleMenusList(c *gin.Context) {
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
	treeMap, err := roleMenusTreeMap()
	var roleMenusList []SysRoleMenusAssignList
	for k, v := range treeMap {
		if k.MenuId == 0 {
			for _, v1 := range v {
				roleMenusList = append(roleMenusList, v1)
			}
		}
	}
	total := len(roleMenusList)
	offset := page.Size * (page.Current - 1)
	limit := page.Size * page.Current
	if total < limit {
		limit = total
	}
	for i := offset; i < limit; i++ {
		sysRoleMenusAssignListChildren(&roleMenusList[i], treeMap)
	}
	_, err = json.Marshal(roleMenusList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedPageList(page, int64(total), roleMenusList, c)
}

func sysRoleMenusAssignListChildren(menu *SysRoleMenusAssignList, treeMap map[*SysRoleMenusAssignKey][]SysRoleMenusAssignList) {
	for k, v := range treeMap {
		if k.RoleId == menu.RoleId && k.MenuId == menu.MenuId {
			for _, v1 := range v {
				menu.Children = append(menu.Children, v1)
			}
		}
	}
	for i := 0; i < len(menu.Children); i++ {
		sysRoleMenusAssignListChildren(&menu.Children[i], treeMap)
	}
}

func RoleMenusAssign(c *gin.Context) {
	var srma SysRoleMenusAssign
	_ = c.ShouldBindJSON(&srma)
	fmt.Println(srma)
	err := roleMenusAssign(srma)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func RoleMenusUpdate(c *gin.Context) {
	var sysRoleMenus SysRoleMenus
	_ = c.ShouldBindJSON(&sysRoleMenus)
	fmt.Println(sysRoleMenus)
	err := roleMenusUpdate(&sysRoleMenus)
	// var err error
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
