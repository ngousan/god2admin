package sysMenu

import (
	"god2admin/public/response"
)

func getMenuChildren(menu *SysMenu, treeMap map[int][]SysMenu) {
	menu.Children = treeMap[int(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		// menu.Children[i].HasChildren = true
		// menu.Children[i].HasChildren = len(treeMap[int(menu.Children[i].ID)]) == 0
		// if !menu.Children[i].HasChildren {
		getMenuChildren(&menu.Children[i], treeMap)
		// }
	}
}

func getMenuDictChildren(dt *response.DictTree, treeMap map[string][]response.DictTree) {
	dt.Children = treeMap[dt.Value]
	for i := 0; i < len(dt.Children); i++ {
		getMenuDictChildren(&dt.Children[i], treeMap)
	}
}
