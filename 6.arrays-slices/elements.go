package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(mySlice[3])

	mySlice[1] = 0

	fmt.Println(mySlice)
	sliceOfslice := mySlice[2:5]
	fmt.Println(sliceOfslice)
}
