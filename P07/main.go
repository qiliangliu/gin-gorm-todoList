package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//1. 创建template文件
	//2. 解析文件
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Println("template parse err: ", err)
		return
	}
	//3. 渲染模板文件
	err = t.ExecuteTemplate(w, "index.tmpl", "小王子")
	if err != nil {
		fmt.Println("reder template failed err: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http start failed err: ", err)
		return
	}
}
