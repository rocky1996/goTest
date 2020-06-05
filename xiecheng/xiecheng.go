package main

import (
	"fmt"
	"net/http"
)

func main()  {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com",
		"http://www.taobao.com",
		"http://www.163.com",
		"http://www.sohu.com",
	}

	c := make(chan string)
	for _,link := range links{
		go checkLink(link,c)
	}

	for i:=0;i<len(links);i++{
		fmt.Println(<-c)
	}
	//
	//fmt.Println(<-c)//等待通道的消息并打印
	//fmt.Println(<-c)//等待通道的消息并打印
	//fmt.Println(<-c)//等待通道的消息并打印
	//fmt.Println(<-c)//等待通道的消息并打印
	//fmt.Println(<-c)//等待通道的消息并打印
	//fmt.Println(<-c)//等待通道的消息并打印
}

func checkLink(link string,c chan string)  {//通道的参数
	_,err := http.Get(link)
	if err != nil{
		fmt.Println(link,"没有连接上")
		c<-"没有连接上"
		return
	}
	fmt.Println(link,"连接上了")
	c<-"连接上了"
}

//go的协程

/*
 * channel通道,就是实现协程之间的通信
   c := make(chan string)
 */
