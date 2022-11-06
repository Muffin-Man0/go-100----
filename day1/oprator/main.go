package main

import (
	"fmt"
)

func test(arr []int)  {
	arr[0] = 3
}
func main()  {
	arr:= []int{1,2,3,5:3}

	// for _,ele := range arr {
	// 	fmt.Println(ele)
	// }
	// arr[0] = 3
	fmt.Println(arr)
	test(arr)
	fmt.Println(arr[0])
	
}