package main

import (
	"fmt"
)

func main()  {
	/*长度为10的数组*/
	var n [10]int
	var i,j int

	var a = [5][2]int{{0,0},{1,2},{2,4},{3,6},{4,8}}

	for i=0;i<10;i++{
		n[i] = i + 100
	}

	for j=0;j<10;j++{
		fmt.Println("Element[%d] = %d\n",j,n[j])
	}

	for i=0;i<5;i++{
		for j=0;j<2;j++{
			fmt.Println("a[%d][%d] = %d\n",i,j,a[i][j])
		}
	}
}
