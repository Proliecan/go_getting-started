package main

import (
	"fmt"
)

func main() {
	var arr [5]int // array of 5 integers
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	fmt.Println("arr[3] =", arr[3])

	arr2 := [5]int{1, 2, 3, 4, 5} // array literal
	fmt.Println(arr2)
}
