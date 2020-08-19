package database

import (
	"god2admin/global"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	dba := global.CONFIG.Mysql
	if db, err := gorm.Open("mysql", dba.Username+":"+dba.Password+"@("+dba.Path+")/"+dba.Dbname+"?"+dba.Config); err != nil {
		global.LOGGER.Error("初始化Mysql数据库失败!", err)
		os.Exit(0)
	} else {
		global.DB = db
		global.DB.DB().SetMaxIdleConns(dba.MaxIdleConns)
		global.DB.DB().SetMaxOpenConns(dba.MaxOpenConns)
		global.DB.LogMode(dba.LogMode)
	}
}
