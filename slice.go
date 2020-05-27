package main

/**
 * len()
   cap()
   append()
   copy()
 */

import (
	"fmt"
	_ "fmt"
)

//func main()  {
//	var numbers []int
//	printSlices(numbers)
//
//
//}
//
//func printSlices(x []int)  {
//	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}

func main()  {
	var numbers = make([]int,3,5)
	printSlices(numbers)
}

func printSlices(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}