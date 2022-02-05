package ceshi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	bt, _ := ioutil.ReadFile("./hello.txt")
	fmt.Fprintln(w, string(bt))
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
