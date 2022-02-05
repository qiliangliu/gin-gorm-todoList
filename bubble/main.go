package main

import (
	"gin_demo/bubble/dao"
	"gin_demo/bubble/routers"
	"gin_demo/bubble/models"
)



func main() {
	//初始化数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	//注册路由
	r := routers.SetupRouter()
	//启动服务
	r.Run()
}
