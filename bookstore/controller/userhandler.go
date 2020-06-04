package controller

import (
	"goTest/bookstore/dao"
	"html/template"
	"net/http"
)

//Login 处理用户登录的函数
func Login(w http.ResponseWriter,r *http.Request)  {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用
	user,_ := dao.CheckUserNameAndPassword(username,password)
	if user != nil{
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,"")
	}else {
		//不正确,暂时去登录页面
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w,"")
	}
}

//Regist 注册函数
func Regist(w http.ResponseWriter,r *http.Request)  {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用
	user,_ := dao.CheckUserNameAndPassword(username,password)
	if user.ID > 0{
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在")
	}else {
		//用户名可用,保存
		dao.SaveUser(username,password,email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"")
	}
}

func CheckUserName(w http.ResponseWriter,r *http.Request)  {
	username := r.PostFormValue("username")
	user,_ := dao.CheckUserName(username)
	if user.ID>0{
		//用户名已存在
		w.Write([]byte("用户名已存在"))
	}else {
		//用户名可用
		w.Write([]byte("用户名可用"))
	}
}
