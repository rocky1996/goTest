package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter,r *http.Request)  {
	//解析模版文件
	//t,_ := template.ParseFiles("index.html")
	t := template.Must(template.ParseFiles("index.html","index2.html"))
	//执行
	//t.Execute(w,"Hello Template")
	t.ExecuteTemplate(w,"index2.html","我要去index2.html里面去")
}

func main()  {
	http.HandleFunc("/testTemplate",testTemplate)
	http.ListenAndServe(":8080",nil)
}
