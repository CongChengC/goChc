package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的多路复用器处理请求!", r.URL.Path)
}

//试用方法 :
//PS D:\goproject\src\goChc\bookstore\webapp> go build main.go
//PS D:\goproject\src\goChc\bookstore\webapp> .\main.exe
//浏览器输入 http://localhost:8080/ 就可以正常访问了。

func main() {
	//创建多路复用器
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080", mux)

}
