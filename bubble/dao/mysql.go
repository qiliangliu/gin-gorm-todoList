package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySql() (err error) {
	//创建数据库
	//sel: CREATE DATABASE bubble
	//连接数据
	dsn := "root:Qq@123456@(101.42.222.42:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := DB.DB()
	return db.Ping()
}
