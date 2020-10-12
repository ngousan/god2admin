package authUserRouters

import (
	"fmt"
	"god2admin/public/request"
	"god2admin/public/response"

	"github.com/gin-gonic/gin"
)

func UserRouterList(c *gin.Context) {
	cc := request.GetClaims(c)
	treeMap, err := routerListTreeMapByUserId(cc.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	routerList := treeMap["/"]
	for k, _ := range routerList {
		for _, v := range treeMap[routerList[k].Group] {
			routerList[k].Children = append(routerList[k].Children, v)
		}
	}
	fmt.Println("routerList", routerList)
	response.OkWithData(routerList, c)
}

// func UserRoutersResp(c *gin.Context) {
// 	userRouters := []AuthUserRouters{
// 		AuthUserRouters{
// 			Path:      "/main",
// 			Redirect:  AuthUserRoutersRedirect{Name: "main"},
// 			Component: "layoutHeaderAside",
// 			Meta:      AuthUserRoutersMeta{Title: "主菜单", Auth: true},
// 			Children: []AuthUserRouters{
// 				AuthUserRouters{
// 					Name:      "main-index",
// 					Path:      "index",
// 					Component: "pages/index",
// 					Meta: AuthUserRoutersMeta{
// 						Title: "首页",
// 						Auth:  true,
// 					},
// 				},
// 				AuthUserRouters{
// 					Name:      "main-sys-user",
// 					Path:      "sys/user",
// 					Component: "pages/sys/user",
// 					Meta: AuthUserRoutersMeta{
// 						Title: "用户管理",
// 						Auth:  true,
// 					},
// 				},
// 				AuthUserRouters{
// 					Name:      "main-sys-role",
// 					Path:      "sys/role",
// 					Component: "pages/sys/role",
// 					Meta: AuthUserRoutersMeta{
// 						Title: "角色管理",
// 						Auth:  true,
// 					},
// 				},
// 				AuthUserRouters{
// 					Name:      "main-sys-menu",
// 					Path:      "sys/menu",
// 					Component: "pages/sys/menu",
// 					Meta: AuthUserRoutersMeta{
// 						Title: "菜单管理",
// 						Auth:  true,
// 					},
// 				},
// 				AuthUserRouters{
// 					Name:      "main-sys-user-roles",
// 					Path:      "sys/user/roles",
// 					Component: "pages/sys/user/roles",
// 					Meta: AuthUserRoutersMeta{
// 						Title: "用户角色分配",
// 						Auth:  true,
// 					},
// 				},
// 				AuthUserRouters{
// 					Name:      "main-sys-role-menus",
// 					Path:      "sys/role/menus",
// 					Component: "pages/sys/role/menus",
// 					Meta: AuthUserRoutersMeta{
// 						Title: "角色菜单分配",
// 						Auth:  true,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	response.OkWithData(userRouters, c)
// }
