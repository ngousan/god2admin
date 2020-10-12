package authUserRouters

import (
	"god2admin/global"
	"strings"
)

func routerListTreeMapByUserId(userId uint) (treeMap map[string][]AuthUserRouters, err error) {
	db := global.DB.Table("sys_user_menus_v")
	var result []AuthUserRoutersView
	treeMap = make(map[string][]AuthUserRouters)
	db.Where("user_id = ?", userId).Scan(&result)
	for _, v := range result {
		var aur AuthUserRouters
		aur.Name = v.Name
		// aur.Path = v.Path
		aur.Component = v.Component
		aur.Redirect.Name = v.Redirect
		aur.Meta.Title = v.Title
		aur.Meta.Auth = v.Auth
		aur.Meta.Cache = v.Cache
		aur.Hidden = v.Hidden
		aur.Group = v.Group
		if v.Redirect != "" || v.Component != "" {
			if v.Name == v.Group {
				aur.Path = "/" + v.Name
				treeMap["/"] = append(treeMap["/"], aur)
			} else {
				aur.Path = strings.Replace(v.Name, v.Group+"-", "", 1)
				aur.Path = strings.Replace(aur.Path, "-", "/", -1)
				treeMap[v.Group] = append(treeMap[v.Group], aur)
			}
		}
	}
	return
}
