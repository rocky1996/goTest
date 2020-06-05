package main

import (
	"fmt"
	_ "fmt"
	"time"
)
/*
 *消费是FIFO,先进先出
*/
//var message = make(chan string,3)

func sample(message chan string)  {
	message <- "Hello goroutine!!!1"
	message <- "Hello goroutine!!!2"
	message <- "Hello goroutine!!!3"
}

func sample2(message chan string)  {
	time.Sleep(2 * time.Second)
	str := <-message
	str = str + "i am goroutine"
	message <- str

	close(message)
}


func main()  {
	var message = make(chan string,3)
	go sample(message)
	go sample2(message)
	time.Sleep(3 * time.Second)
	//fmt.Println(<-message)
	for str:= range message{
		fmt.Println(str)
	}

	fmt.Println("hello world")
}
