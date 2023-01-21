package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:] // slice of arr

	arr[1] = 7
	slice[2] = 8

	fmt.Println(arr, slice) // slice and array share the same underlying array

	slice2 := arr[1:3] // slice of arr from index 1 to 3
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)

	slice3 = append(slice3, 6, 7, 8)
	fmt.Println(slice3)

	s1 := slice3[1:]
	s2 := slice3[:2]
	s3 := slice3[1:2]
	fmt.Println(s1, s2, s3)
}
