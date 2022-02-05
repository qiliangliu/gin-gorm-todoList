package main

import (
	"fmt"
	"gorm.io/driver/mysql"
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
	//db.AutoMigrate(&User{})
	//user := User{}
	//res := db.Last(&user)
	//time.Sleep(5 * time.Second)
	//res = db.First(&user)
	//fmt.Println(user)
	//fmt.Println(errors.Is(res.Error, gorm.ErrRecordNotFound))
	//var users []User
	//db.Find(&users)
	//for _, u := range users {
	//	fmt.Println(u)
	//}
	//var user User
	//db.Find(&users, []int{5, 6, 7, 8, 9, 10})
	//fmt.Println(users)

	//var user User
	//db.First(&user)
	//fmt.Println(user)
	//db.Take(&user)
	//db.Last(&user)
	//db.Find(&users)
	//fmt.Println(users)
	//db.First(&user, 10)
	//db.Where("name = ?", "dog2").First(&user)
	//db.Where("age = ?", 18).Find(&users)
	//db.Where("name <> ?", "dog2").Find(&users)
	//db.Where("name IN (?)", []string{"dog1", "dog2"}).Find(&users)
	//db.Where("name LIKE ?", "%dog%").Find(&users)
	//db.Where("name = ? AND age >= ?", "启亮", "18").Find(&users)
	//db.Where("updated_at > ?", "2000-01-01 00:00:00").Find(&users)
	//db.Where("age BETWEEN ? AND ?", 10, 20).Find(&users)
	//db.Where(&User{ Age: 18 }).Find(&users)
	//db.Where(map[string]interface{} {"age": 28}).Find(&users)
	//db.Where([]int64{ 1, 2, 3 }).Find(&users)
	//db.Not("name", "dog1").Find(&users)
	//db.Not("name", []string{ "dog1", "dog2" }).Find(&users)
	//db.Not([]int64{ 1, 2, 3 }).Find(&users)
	//db.Not([]int64{}).Find(&users)
	//db.Not("age = ?",  18).Find(&users)
	//db.Where("name = ?", "dog1").Or("age = ?", 18).Find(&users)
	//db.Where("name = 'dog1'").Or(&User{ Name: "dog2"}).Find(&users)
	//db.First(&user, 10)
	//db.Find(&user, "name = ?", "dog1")
	//db.Find(&users, "name <> ? AND age > ?", "dog1", 18)
	//db.Find(&users, User{Age: 38})
	//db.Find(&users, map[string]interface{} { "age": 48 })
	//db.FirstOrInit(&user, User{ Name: "none"})
	//db.Where(User{ Name: "dog2"}).FirstOrInit(&user)
	//db.FirstOrInit(&user, map[string]interface{} { "Name": "dog2" })
	//db.Where(User{ Name: "none" }).Attrs(User{ Age: 20 }).FirstOrInit(&user)
	//db.Where(User{ Name: "dog2"}).Assign(User{ Age: 66}).FirstOrInit(&user)
	//db.FirstOrCreate(&user, User{ Name: "not_exist"})
	//db.Where(User{ Name: "none" }).Attrs(&User{Age: 20}).FirstOrCreate(&user)
	//db.Where(User{ Name: "none"}).Assign(&User{Age: 20}).FirstOrCreate(&user)
	//var user User
	//var users []User
	//var users1 []User
	//var users2 []User
	//db.Select("name, age").Find(&users)
	//db.Select([]string{ "name", "age" }).Find(&users)
	//db.Select("*").Find(&users)
	//db.Table("users").Select("COALESCE(age, ?)", -1).Find(&users)
	//db.Order("age desc, name asc").Find(&users)
	//db.Order("age desc").Order("name").Find(&users)
	//db.Debug().Order("age desc").Find(&users1).Order("age").Find(&users2)
	//for _, v := range users1 {
	//	fmt.Println(v.Age)
	//}
	//fmt.Println();
	//for _, v := range users2 {
	//	fmt.Println(v.Age)
	//}
	//fmt.Println(users1)
	//fmt.Println(users2)

	//var users1 []User
	//var users2 []User
	//db.Limit(3).Find(&users)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	//fmt.Println(users1)
	//fmt.Println("\n\n\n")
	//fmt.Println(users2)
	//db.Offset(2).Find(&users)
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2);
	//var count int64;
	//var users []User
	//db.Where("name = ?", "dog1").Or("name = ?", "dog2").Find(&users).Count(&count)
	//db.Model(&User{}).Where("name = ?", "dog1").Count(&count)
	//db.Table("users").Count(&count)
	//fmt.Println(users)
	//fmt.Println(count)

	//rows, _ := db.Debug().Table("users").Select("'name', age").Group("age").Rows()
	//for rows.Next() {
	//	s, _ := rows.Columns()
	//	fmt.Println(s)
	//}

	//var users []User
	//db.Table("users").Select("'name', age").Group("age").Having("age > ?", 18).Scan(&users)
	//fmt.Println(users)

	//rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.name = user.name").Rows()
	//db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.name = users.name").Scan(&users)
	//db.Joins("JOIN emails ON emails.name = users.name AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.name = users.name").Where("credit_cards.number = ?", "4111111111").Find(&users)

	//var ages []int64
	//var users []User
	//db.Debug().Model(User{}).Pluck("Age", &ages)
	//fmt.Println(ages)

	//tx := db.Where("name = ?", "lql")
	//if someCondition {
	//	tx = tx.Where("age = ?", 20)
	//}
	//else {
	//	tx = tx.Where("age = ?", 30)
	//}
	//
	//if yetAnotherCondition {
	//	tx = tx.Where("active = ?", 1)
	//}

	//var user User
	//db.First(&user)
	//user.Name = "七米"
	//user.Age = 99
	//db.Save(&user)
	//db.Model(&User{}).Where("age = ?", 18).Update("age", 99)
	//db.Model(&user).Updates(map[string]interface{} { "name": "lql", "age": 18 })
	//db.Model(&user).Omit("name").Updates(map[string]interface{} { "name": "lql666", "age": 18 })
	//db.Table("users").Where("id IN (?)", []int{ 10, 11 }).Updates(map[string]interface{} { "name": "hello" })
	//fmt.Println(user)
	var user User
	//db.First(&user)
	//user.ID = 10
	//db.Delete(&user)
	//fmt.Println(db.Model(User{}).Where("id = ?", 10).Updates(User{ Name: "lazy_cat", Age: 4 }).RowsAffected)
	//db.Where("name LIKE ?", "%dog%").Delete(&User{})
	//db.Delete(&User{}, "name LIKE ?", "%hell%")
	//db.Unscoped().Where("name = ?", "dog2").First(&user)
	db.Unscoped().Where("age = ?", 18).Delete(&User{})
	fmt.Println(user)
}
