package main

import (
	"goTest/bookstore/controller"
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

	//去登录
	http.HandleFunc("/login",controller.Login)
	//去注册
	http.HandleFunc("/Regist",controller.Regist)
	//去注册
	http.HandleFunc("/CheckUserName",controller.CheckUserName)
	//获取所有图书
	http.HandleFunc("/GetBooks",controller.GetBooks)
	//删除图书
	http.HandleFunc("/DeleteBooks",controller.DeleteBooks)
	//根据id去查找图书
	http.HandleFunc("/ToUpdateBookPage",controller.ToUpdateBookPage)
	//更新图书
	http.HandleFunc("/UpdateBook",controller.UpdateBook)


	http.ListenAndServe(":8080",nil)
}
