package main

import (
	"html/template"
	"net/http"
)

// 创建处理器函数
// http://localhost:8080/testTemplate?
func testTemplate(w http.ResponseWriter, r *http.Request) {
	//例子1
	// //解析模板文件
	// t, _ := template.ParseFiles("index.html")
	// //执行
	// t.Execute(w, "Hello Template")

	//例子2
	//通过Must函数让Go帮我们自动处理异常
	t := template.Must(template.ParseFiles("index.html", "index2.html"))
	//将响应数据在index2.html文件中显示
	t.ExecuteTemplate(w, "index2.html", "我要去index2.html中")
}

func main() {

	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
