package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("newname", "./statics")
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//1 创建模板
	//2 解析模板
	r.LoadHTMLGlob("all/**/*")
	//3 渲染模板
	r.GET("/tem_index", func(c *gin.Context) {
		//返回状态信息
				//状态码        渲染的那个文件			渲染的内容
		c.HTML(http.StatusOK, "templates_index.templ", gin.H{
			"title": "<hr>",
		})
	})

	r.GET("/pos_index", func(c *gin.Context) {
		//返回状态信息
		//状态码        渲染的那个文件			渲染的内容
		c.HTML(http.StatusOK, "posts_index.templ", gin.H{
			"title": "pos_index",
		})
	})

	r.Run(":9090")
}
