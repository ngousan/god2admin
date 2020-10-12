package sysMenu

import (
	"encoding/json"
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	page, err := request.GetPageListInfo(c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	var menuList []SysMenu
	treeMap, err := listTreeMap()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	menuList = treeMap[0]
	total := len(menuList)
	offset := page.Size * (page.Current - 1)
	limit := page.Size * page.Current
	if total < limit {
		limit = total
	}

	for i := offset; i < limit; i++ {
		// menuList[i].HasChildren = true
		// menuList[i].HasChildren = len(treeMap[int(menuList[i].ID)]) == 0
		// if !menuList[i].HasChildren {
		getMenuChildren(&menuList[i], treeMap)
		// }
	}
	_, err = json.Marshal(menuList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	response.SucceedPageListTree(page, int64(total), menuList, c)
}

func Create(c *gin.Context) {
	var sysMenu SysMenu
	_ = c.ShouldBindJSON(&sysMenu)
	err := insert(sysMenu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func Update(c *gin.Context) {
	var sysMenu SysMenu
	_ = c.ShouldBindJSON(&sysMenu)
	err := update(&sysMenu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// 菜单字典
func MenuDict(c *gin.Context) {
	treeMap, err := dictTreeMap()
	menuList := treeMap["0"]
	// for _, v := range menuList {
	// 	getMenuDictChildren(&v, treeMap)
	// }
	for i := 0; i < len(menuList); i++ {
		getMenuDictChildren(&menuList[i], treeMap)
	}
	fmt.Println(menuList)
	_, err = json.Marshal(menuList)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取菜单字典失败，%v", err), c)
		return
	}
	response.SucceedDict(menuList, c)
}
