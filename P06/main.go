package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func myfun(w http.ResponseWriter, r *http.Request) {
	//自定义一个函数kua
	kua := func(name string) (string, error) {
		return name + "又年轻又帅！", nil
	}

	//定义模板
	t := template.New("f.tmpl")	//创建一个名字是f.tmpl的模板对象t, 注意这个必须与所解析的文件的名字一样

	//告诉模板引擎，我现在多了一个自定了函数kua
	t.Funcs(template.FuncMap{
		"kua": kua,
	})

	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("parse tempalte failed, err: ", err);
		return
	}

	//渲染模板
	err = t.Execute(w, "小王子")
	if err != nil {
		fmt.Println("template execute err: ", err)
	}
}

func mytem(w http.ResponseWriter, r *http.Request) {
	//创建模板
	//解析模板
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Println("template parse failed, err: ", err)
		return
	}
	//渲染模板
	t.Execute(w, "小王子")
}

func main() {
	http.HandleFunc("/", myfun)
	http.HandleFunc("/template_demo", mytem)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http start failed, err: ", err)
	}
}
