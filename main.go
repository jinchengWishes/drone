package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/kotlin2018/adapter"
)

func main() {
	opts := &adapter.Adapter{
		DriverName: "mysql",
		LinkInfo:   "root:root@tcp(127.0.0.1:3306)/golang",
		TableName:  "casbin_rule",
	}

	c := &adapter.CasBinModel{
		BaseAdapter: opts,
		ModelPath:   "./rbac_model.conf",
	}

	e, err := casbin.NewEnforcer(c, opts)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("ok:", e)

	// boolean, err := e.Enforce("100", "/v1/post/info", "post")
	// if err != nil {
	// 	return
	// }
	// if boolean {
	// 	fmt.Println("有权限")
	// } else {
	// 	fmt.Println("没有权限")
	// }
}

// http server

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, map[interface{}]interface{}{
// 		"name": "jincheng",
// 		"age":  "20",
// 	})
// }
