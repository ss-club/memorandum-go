package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		println("数据库连接错误")
		panic(err)
	}
	println("数据库连接成功")
	DB = db
	migration()
}
