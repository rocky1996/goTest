package main

import (
	"fmt"
	"strconv"
	"time"
)

func sample3(ch chan string)  {
	for i:=0;i<19;i++{
		ch <- "I am sample1 num:"+strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
}

func sample4(ch chan int)  {
	for i:=0;i<10;i++{
		ch <- i
		time.Sleep(2 * time.Second)
	}
}

func main()  {
	ch3 := make(chan string,3)
	ch4 := make(chan int,5)

	for i:=0;i<10;i++{
		go sample3(ch3)
		go sample4(ch4)
	}

	for i:=0;i<1000;i++{
		select {
		case str,ch1Check := <-ch3:
			if !ch1Check {
				fmt.Println("ch1Check false")
			}
			fmt.Println(str)
		case p,ch2Check := <-ch4:
			if !ch2Check {
				fmt.Println("ch2Check false")
			}
			fmt.Println(p)
		}
	}
}
