package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器搜索栏中返回的 query string 参数
		//name := c.Query("query")
		//name := c.DefaultQuery("query", "nobody")
		name, ok := c.GetQuery("query")
		age := c.Query("age")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"error": "no parameter",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": name,
				"age": age,
			})
		}
	})

	r.Run(":9090")
}
