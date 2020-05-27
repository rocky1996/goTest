package main

import (
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {
	fmt.Println(Books{"Go 语言","www.runoob.com","Go语言教程",5495470})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}