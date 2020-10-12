package ProviderSysMenu

import menu "god2admin/module/sys/menu"

func PRVDuserMenu() (userMenus []menu.SysMenu) {
	// 获取userId
	// 获取userRoles，RoleId
	// 获取RoleMenus，MenuId
	// 获取UserMenus
	userMenus = []menu.SysMenu{
		menu.SysMenu{
			Title: "主菜单",
			Icon:  "home",
			Path:  "",
			Children: []menu.SysMenu{
				menu.SysMenu{
					Title: "首页",
					Icon:  "home",
					Path:  "/main/index",
				},
				menu.SysMenu{
					Title: "系统管理",
					Icon:  "cog",
					Path:  "",
					Children: []menu.SysMenu{
						menu.SysMenu{
							Title: "用户管理",
							Icon:  "cog",
							Path:  "/main/sys/user",
						},
						menu.SysMenu{
							Title: "菜单管理",
							Icon:  "cog",
							Path:  "/main/sys/menu",
						},
						menu.SysMenu{
							Title: "角色管理",
							Icon:  "cog",
							Path:  "/main/sys/role",
						},
					},
				},
			},
		},
	}
	return
}
