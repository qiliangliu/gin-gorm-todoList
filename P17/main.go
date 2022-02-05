package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{ "method": "GET" })
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{ "method": "POST" })
			//...
		}
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{ "msg": "你访问的页面不存在啊！！！" })
	})

	//路由组的组
	//把公共前缀提取出来，创建一个路由组
	//路由组可以嵌套使用
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/kind", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{ "msg": "/video/kind" })
		})
		kindGroup := videoGroup.Group("/kind")
		{
			kindGroup.GET("/tragedy", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{ "msg": "/video/kind/tragedy" })
			})
			kindGroup.GET("/comedy", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{ "msg": "/video/kind/comedy" })
			})
		}
	}

	r.Run(":9090")
}
