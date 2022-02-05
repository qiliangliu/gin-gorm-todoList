package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func attack(w http.ResponseWriter, r *http.Request) {
	//1. 创建模板文件
	//2. 解析模板文件
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err: ", err)
		return
	}
	//3. 渲染模板文件
	js1 := "<script>alert('123')</script>"
	js2 := "<a href='https//liwenzhou.com'> LWZ的博客</a>"
	err = t.Execute(w, map[string]string{
		"js1": js1,
		"js2": js2,
	})
}


func main() {
	http.HandleFunc("/attack", attack)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http start failed err: ", err)
		return
	}
}
