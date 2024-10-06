package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "通过自己创建的处理器处理请求！")
	fmt.Fprintln(w, "通过详细配置服务器的信息来处理请求！")
}

func main() {

	myHandler := MyHandler{}

	//例子1
	http.Handle("/myHandler", &myHandler)
	http.ListenAndServe(":8080", nil)

	//创建Server结构，并详细配置里面的字段
	// server := http.Server{
	// 	Addr:        ":8080",
	// 	Handler:     &myHandler,
	// 	ReadTimeout: 2 * time.Second,
	// }
	// server.ListenAndServe()

}