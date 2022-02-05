package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//1. 创建模板，在hello.tmpl中完成了
func sayHello(w http.ResponseWriter, r *http.Request) {
	//2. 解析所创建的模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed err: ", err)
		return
	}
	//3. 渲染模板
	err = t.Execute(w, "小王子")
	if err != nil {
		fmt.Println("render template failed err: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed, err: ", err)
		return
	}
}