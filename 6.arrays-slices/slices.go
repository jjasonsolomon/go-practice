package main

import (
	"fmt"
)

func main() {

	fruits := []string{"Apple", "Orange", "Mango"}

	fmt.Println("length of slice is %d and capacity is %d\n", len(fruits), cap(fruits))

	for _, i := range fruits {

		fmt.Println(i)
	}
}
