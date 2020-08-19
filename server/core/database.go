package core

import (
	"god2admin/database"
	"god2admin/global"
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
