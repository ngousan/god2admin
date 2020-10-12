package database

import (
	"god2admin/global"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

// 初始化数据库并产生数据库全局变量
// func Mysql() {
// 	dba := global.CONFIG.Mysql
// 	if db, err := gorm.Open("mysql", dba.Username+":"+dba.Password+"@("+dba.Path+")/"+dba.Dbname+"?"+dba.Config); err != nil {
// 		global.LOGGER.Error("初始化Mysql数据库失败!", err)
// 		os.Exit(0)
// 	} else {
// 		global.DB = db
// 		global.DB.DB().SetMaxIdleConns(dba.MaxIdleConns)
// 		global.DB.DB().SetMaxOpenConns(dba.MaxOpenConns)
// 		global.DB.LogMode(dba.LogMode)
// 	}
// }

// GormMysql 初始化Mysql数据库
func Mysql() {
	dba := global.CONFIG.Mysql
	dsn := dba.Username + ":" + dba.Password + "@tcp(" + dba.Path + ")/" + dba.Dbname + "?" + dba.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	// gormConfig := config(dba.LogMode)
	if global.DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}); err != nil {
		global.LOGGER.Error("初始化Mysql数据库失败!", err)
		os.Exit(0)
	} else {
		// GormDBTables(global.DB)
		db, _ := global.DB.DB()
		db.SetMaxIdleConns(dba.MaxIdleConns)
		db.SetMaxOpenConns(dba.MaxOpenConns)
	}
}
