package routers

import (
	"gin_demo/bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//告诉我们的gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//加载要html文件
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看
		//查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
