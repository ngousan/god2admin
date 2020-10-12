package core

import (
	"god2admin/database"
	"god2admin/global"
	sysMenu "god2admin/module/sys/menu"
	sysRole "god2admin/module/sys/role"
	sysUser "god2admin/module/sys/user"
	sysUserRoles "god2admin/module/sys/user/roles"
)

// Initial database
func initDatabase() {
	switch global.CONFIG.System.DbType {
	case "mysql":
		database.Mysql()
	// case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		database.Mysql()
	}
}

// 注册数据库表专用
func registerTables() {
	db := global.DB
	db.AutoMigrate(
		sysUser.SysUser{},
		sysUserRoles.SysUserRoles{},
		sysMenu.SysMenu{},
		sysRole.SysRole{},
	)
	global.LOGGER.Debug("register tables successful")
}
