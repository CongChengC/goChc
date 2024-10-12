package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "测试http协议")
}

// 测试:
// http://localhost:8080/http
// 浏览器中F12--》network--》刷新--》点弹出的http链接--》点Headers就可以查看
// 用index.html 中的前端代码进行访问，360浏览器中顺利看到了form data;但谷歌gpt浏览器没有看到，可能哪里设置被拦截了吗？？
func main() {
	//调用处理器请求
	http.HandleFunc("/http", handler)
	//路由
	http.ListenAndServe(":8080", nil)
}
