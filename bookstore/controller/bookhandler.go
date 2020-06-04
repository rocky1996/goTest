package controller

import (
	"goTest/bookstore/dao"
	"goTest/bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

//GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter,r *http.Request)  {
	books,_ := dao.GetBooks()
	//解析模版文件
	t := template.Must(template.ParseFiles("/views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}

//DeleteBooks 删除图书
func DeleteBooks(w http.ResponseWriter,r *http.Request)  {
	bookId := r.PostFormValue("bookId")
	dao.DeleteBook(bookId)
	//解析模版文件
	t := template.Must(template.ParseFiles("/views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w,"")
}

//根据id去查找
func ToUpdateBookPage(w http.ResponseWriter,r *http.Request)  {
	bookId := r.PostFormValue("bookId")
	book,_ := dao.GetBookById(bookId)
	t := template.Must(template.ParseFiles("/views/pages/manager/book_modify.html"))
	t.Execute(w, book)
}

func UpdateBook(w http.ResponseWriter,r *http.Request)  {
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	//将价格，销量和库存进行转化
	fPrice,_ := strconv.ParseFloat(price,64)
	iSales,_ := strconv.ParseInt(sales,10,0)
	iStock,_ := strconv.ParseInt(stock,10,0)
	iBookId,_ := strconv.ParseInt(bookId,10,0)

	book := &model.Book{
		ID:       int(iBookId),
		Title:    title,
		Author:   author,
		Price:    fPrice,
		Sales:    int(iSales),
		Stock:    int(iStock),
		ImgPath:  "/dedeed",
	}

	dao.UpdateBook(book)
	GetBooks(w,r)
}