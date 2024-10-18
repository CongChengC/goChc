package main

import (
	"fmt"
	"net/http"
)

// setCookie 添加Cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	//创建Cookie
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: true,
	}
	//将Cookie发送给浏览器
	// w.Header().Set("Set-Cookie", cookie.String())
	//添加第二个Cookie
	// w.Header().Add("Set-Cookie", cookie2.String())
	//直接调用http的SetCookie函数设置Cookie
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

// getCookies 获取Cookie
func getCookies(w http.ResponseWriter, r *http.Request) {
	//获取请求头中所有的Cookie
	// cookies := r.Header["Cookie"]
	// fmt.Println("得到的Cookie有：", cookies)

	//如果想得到某一个Cookie，可以直接调用Cookie方法
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的Cookie有：", cookie)

}

func main() {
	//http://localhost:8080/setCookie;调用后打开F12就可以查看到
	http.HandleFunc("/setCookie", setCookie)
	//http://localhost:8080/getCookies
	http.HandleFunc("/getCookies", getCookies)

	http.ListenAndServe(":8080", nil)
}
