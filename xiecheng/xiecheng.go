package main

import (
	"fmt"
	"time"
)

//总票数
var tickCount = 200

/**
 * 购票
 */
func buy()  {
	tickCount = tickCount-1
}

/**
 * 退票
 */
func refund()  {
	tickCount = tickCount+1
}

func main()  {
	// 购票协程
	go refund()
	// 退票协程
	go buy()

	// 为了简单，这里睡眠1秒，等待上面两个协程结束
	time.Sleep(time.Second * 1)

	fmt.Println(tickCount) // 输出结果是什么？
}