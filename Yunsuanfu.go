package main

import (
	"fmt"
)

func main()  {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Println("第一行 - c的值为%d\n",c)

	a++
	fmt.Println("第一行 - a的值为%d\n",a)

	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )
}