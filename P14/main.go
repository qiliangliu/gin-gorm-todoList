package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username"`
	Age      int    `form:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/form", func(c *gin.Context) {
		u := User{
			"",
			-1,
		}
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": u.Username,
				"age":  u.Age,
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		u := User{
			"",
			-1,
		}
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": u.Username + "666",
				"age":  u.Age,
			})
		}
	})

	r.POST("/jason", func(c *gin.Context) {
		u := User{
			"",
			-1,
		}
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": u.Username + "888",
				"age":  u.Age,
			})
		}
	})
	r.Run(":9090")
}
