package main

import (
	"fmt"
)

func main()  {
	//var a int = 100
	//var b int = 200
	//var ret int
	//
	//ret = max(a,b)
	//fmt.Println("最大值是:%d\n",ret)

	a,b := swap("Google","Runoob")
	fmt.Println(a,b)
}

/**
 * 定义方法
 */
func max(num1, num2 int) int {
	//定义局部变量
	var result int
	if(num1 > num2){
		result = num1
	}else {
		result = num2
	}
	return result
}

/**
 * 参数交换方法
 */
func swap(x,y string) (string,string) {
	return y,x
}
