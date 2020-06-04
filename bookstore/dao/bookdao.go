package dao

import (
	"goTest/bookstore/model"
	"goTest/bookstore/utils"
)

//获取全部图书
func GetBooks() ([]*model.Book,error)  {
	sqlStr := "select id,title,author,price,sales,stock,imgPath from books"
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		return nil,err
	}

	var books []*model.Book

	for rows.Next(){
		var book *model.Book
		//给book中的字段赋值
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books = append(books,book)
	}
	return books,nil
}

//删除图书
func DeleteBook(bookId string) (error) {
	sqlStr := "delete from books where id = ?"

	_, err := utils.Db.Exec(sqlStr,bookId)
	if err != nil{
		return err
	}
	return nil
}

//GetBookById,根据id获取图书
func GetBookById(bookId string) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,imgPath from books where id = ?"
	row := utils.Db.QueryRow(sqlStr,bookId)

	book := &model.Book{}
	row.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
	return book,nil
}

//更新图书信息
func UpdateBook(book *model.Book) (error) {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=?,imgPath=? where id = ?"
	_, err := utils.Db.Exec(sqlStr,book.Title,book.Author,book.Price,book.Sales,book.Stock,book.ImgPath,book.ID)
	if err != nil{
		return err
	}
	return nil
}
