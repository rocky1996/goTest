package main

import (
	json2 "encoding/json"
	"goTest/web/model"
	"net/http"
)

//创建处理器函数
//func handler(w http.ResponseWriter, r *http.Request)  {
//	fmt.Fprintln(w,"你发送的请求的请求地址是:",r.URL.Path)
//	fmt.Fprintln(w,"你发送的请求的请求地址后的查询字符串是:",r.URL.RawQuery)
//	fmt.Fprintln(w,"你发送的请求头的所有信息:",r.Header)
//	fmt.Fprintln(w,"你发送的请求头的Accept-Encoding的信息:",r.Header["Accept-Encoding"])
//	fmt.Fprintln(w,"你发送的请求头的Accept-Encoding的属性值:",r.Header.Get("Accept-Encoding"))
//
//	//获取请求体内容的长度
//	len := r.ContentLength
//	//创建切片
//	body := make([]byte,len)
//	//将请求体内容读到body里面
//	r.Body.Read(body)
//	//在浏览器里面显示请求体中的内容
//	fmt.Fprintln(w,"请求体中的内容是:",string(body)
//
//	//解析表单，在调用之前
//	r.ParseForm()
//	//获取请求参数
//	fmt.Fprintln(w,"请求参数有:",r.Form)
//}

func testJsonRes(w http.ResponseWriter,r *http.Request)  {
	//设置响应内容的类型
	w.Header().Set("Content-Type","application/json")
	//创建User
	user := model.User{
		ID:1,
		Username: "admin",
		Password: "123456",
		Email: "fneucrf@qq.com",
	}
	//将user转换为json格式
	json,_ := json2.Marshal(user)
	//将json格式响应给客户端
	w.Write(json)
}

func testRedir(w http.ResponseWriter,r *http.Request)  {
	//设置响应内的Location属性
	w.Header().Set("Location","http//www.baidu.com")
	//设置响应的状态码
	w.WriteHeader(302)
}

func main()  {
	//http.HandleFunc("/hello",handler)
	//http.HandleFunc("/testJson",testJsonRes)
	http.HandleFunc("/testRedir",testRedir)
	http.ListenAndServe(":8080",nil)
}
