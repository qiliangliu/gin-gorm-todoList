package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := gin.H{ "name": "小王子", "message": "hello world!", "age": 18 }
		c.JSON(http.StatusOK, data)
	})

	type msg struct {
		Name string `json:"name"`
		Message string `json:"message"`
		Age int `json:"age"`
	}
	r.GET("/struct_json", func(c *gin.Context) {
		data := msg{
			"小王子",
			"Helow Golang!",
			18,
		}
		c.JSON(http.StatusOK, data)
	})


	r.Run(":9090")
}
