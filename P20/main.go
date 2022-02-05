package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name string
	Age int64
}

type Bnimal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name string `gorm:"column:my_name"`
	Age int64
}

//自定义创建Animal表时候的名字
func (Animal) TableName() string {
	return "QiLiang"
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "my_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:Qq@123456@(101.42.222.42:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)		//禁用命名表名的时候自动添加s

	//创建表 自动迁移(把结构体和数据库进行对应，创建新表)
	//db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := UserInfo{1, "小王子", "男", "蛙泳"}
	//u2 := UserInfo{2, "lql", "男", "ball"}
	//db.Create(&u1)

	//查询数据
	//var u UserInfo
	//db.Debug().First(&u)
	//fmt.Println(u)

	//db.Model(&u).Update("hobby", "双色球")

	//删除数据
	//db.Delete(&u)
	//db.AutoMigrate(&Animal{})
	//db.Table("用户信息").CreateTable(&UserInfo{})
	db.AutoMigrate(&Bnimal{})
}
