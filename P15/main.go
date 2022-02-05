package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		f, err := c.FormFile("f1")
		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "bard error",
			})
			return
		}

		dfs := path.Join("./", f.Filename)
		c.SaveUploadedFile(f, dfs)
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":9090")
}
