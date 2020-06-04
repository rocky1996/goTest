package web_cookie

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter,r *http.Request)  {
	//创建Cookie
	cookie := http.Cookie{
		Name:"user",
		Value:"admin",
		HttpOnly:true,
	}
	cookie2:= http.Cookie{
		Name:"user",
		Value:"admin",
		HttpOnly:true,
	}
	////将Cookie发送给刘篮球
	//w.Header().Set("Set-Cookie",cookie.String())
	//
	////添加第二个cookie
	//w.Header().Add("Set-Cookie",cookie2.String())
	http.SetCookie(w,&cookie)
	http.SetCookie(w,&cookie2)
}

func getCookies(w http.ResponseWriter,r *http.Request)  {
	//获取请求头里面所有的cookie
	cookies := r.Header["Cookie"]
	fmt.Println("得到的Cookie有:",cookies)
}

func main()  {
	http.HandleFunc("/setCookie",setCookie)
	http.ListenAndServe(":8080",nil)
}
