package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//username := c.DefaultPostForm("username", "0.0")
		//password := c.DefaultPostForm("password", "null")
		username, ok := c.GetPostForm("username")
		if !ok {
			return
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
