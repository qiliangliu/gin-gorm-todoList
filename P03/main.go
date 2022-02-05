package main

import (
	"github.com/gin-gonic/gin"
)

func sayHellow(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hellow world!",
	})
}


func main() {
	r := gin.Default()
	r.GET("/hello", sayHellow)
	r.Run(":9090")
}

