package main

import (
	"fmt"
)

type Book struct{
	title string
	author string
	subject string
	book_id int
}

func main()  {
	var book1 Book        /* Declare Book1 of type Book */
	var book2 Book        /* Declare Book2 of type Book */

	/* book 1 描述 */
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407

	/* book 2 描述 */
	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.book_id = 6495700

	printBook(&book1)
	printBook(&book2)
}

func printBook(book *Book)  {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}