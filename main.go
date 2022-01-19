package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("启动服务成功...")
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, map[interface{}]interface{}{
		"name": "jincheng",
		"age":  "20",
	})
}
