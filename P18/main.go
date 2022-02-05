package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	c.Next()
	//c.Abort() //阻止执行后面的函数，但是这个函数会执行到末尾
	cost := time.Since(start)
	fmt.Println("cost: ", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "lql")		//注意Set与Next直接的调用顺序，对name赋值之间的区别
	c.Next()
	c.Copy()
	fmt.Println("m2 out...")
}

//应用举例，通过闭包的形式，用中间件的方式用户登录验证操作
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其他的准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if 是登录用户
			//c. Next()
			//else
			//c.Abort()

		} else {
			c.Next() //如果需要中间件检查，则直接放行
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(false))
	r.GET("/index", func(c *gin.Context) {
		name, ok := c.Get("name")
		if !ok {
			name = "匿名用户"
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": name,
		})
	})

	//为路由组创建中间件
	//1. 直接在创建路由表的时候，写入中间件函数
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "msg": "xxGroup" })
		})
	}

	//2. 使用路由组函数成员函数Use()
	xxGroup2 := r.Group("/xx2")
	xxGroup2.Use(authMiddleware(true))
	{
		xxGroup2.GET("index2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{ "msg": "xxGroup2" })
		})
	}


	r.Run(":9090")
}
