package main

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	//1. 连接数据库
	dsn := "root:Qq@123456@(101.42.222.42:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})


	//var myuser = []User{
	//	{ Name: "dog1"},
	//	{ Name: "dog2"},
	//	{ Name: "dog3"},
	//}
	//db.Create(&myuser)
	//for _, u := range myuser {
	//	fmt.Println(u.ID)
	//}
	//var users = []User{{Name: "jinzhu1", Birthday: time.Now()}, {Name: "jinzhu2", Birthday: time.Now()}, {Name: "jinzhu3", Birthday: time.Now()}}
	//db.Create(&users)
	//
	//for _, u := range users {
	//	fmt.Println(u.ID)
	//}

	//var users = []User{
	//	{Name: "jinzhu_1"},
	//	{Name: "jinzhu_2"},
	//	{Name: "jinzhu_3"},
	//	{Name: "jinzhu_4"},
	//	{Name: "jinzhu_5"},
	//	{Name: "jinzhu_6"},
	//	{Name: "jinzhu_7"},
	//	{Name: "jinzhu_8"},
	//	{Name: "jinzhu_9"},
	//	{Name: "jinzhu_10"},
	//}
	//
	//// 数量为 100
	//db.CreateInBatches(users, 10)

	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "启亮", "Age": 18,
	})
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "张1", "Age": 28},
	//	{"Name": "张2", "Age": 38},
	//	{"Name": "张3", "Age": 48},
	//	{"Name": "张4", "Age": 58},
	//})

}
