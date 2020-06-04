package main

import (
	"html/template"
	"net/http"
)

//去首页
func IndexHandler(w http.ResponseWriter,r *http.Request){
	//解析模版
	t := template.Must(template.ParseFiles("index.html"))
	//执行
	t.Execute(w,"")
}

func main()  {

	//设置处理静态资源,css,js
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("/views/static/"))))

	//直接去html页面
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("/views/pages/"))))

	http.HandleFunc("/main",IndexHandler)
	http.ListenAndServe(":8080",nil)
}
