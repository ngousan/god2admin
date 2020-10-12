package authUserMenus

import (
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"

	"github.com/gin-gonic/gin"
)

func authUserMenusChildren(menu *AuthUserMenus, treeMap map[int][]AuthUserMenus, parentPath string) {
	menu.Children = treeMap[int(menu.MenuId)]
	if menu.ParentId != 0 {
		menu.Path = parentPath + "/" + menu.Path
	}
	for i := 0; i < len(menu.Children); i++ {
		authUserMenusChildren(&menu.Children[i], treeMap, menu.Path)
	}
}

func UserMenuList(c *gin.Context) {
	cc := request.GetClaims(c)
	treeMap, err := menuListTreeMapByUserId(cc.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	menuList := treeMap[0]
	for i := 0; i < len(menuList); i++ {
		authUserMenusChildren(&menuList[i], treeMap, menuList[i].Path)
	}
	response.OkWithData(menuList, c)
}
