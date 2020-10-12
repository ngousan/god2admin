package authUserMenus

import (
	"god2admin/global"
)

func menuListTreeMapByUserId(userId uint) (treeMap map[int][]AuthUserMenus, err error) {
	db := global.DB.Table("sys_user_menus_v")
	var result []AuthUserMenus
	treeMap = make(map[int][]AuthUserMenus)
	db.Where("user_id = ?", userId).Scan(&result)
	for _, v := range result {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return
}
