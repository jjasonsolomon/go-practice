package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	for _, i := range mySlice {

		fmt.Println(i)
	}

	newSlice := []int{10, 20, 30}

	mySlice = append(mySlice, newSlice...)

	fmt.Println("Slice length is %d but capacity is %d\n", len(mySlice), cap(mySlice))

}
