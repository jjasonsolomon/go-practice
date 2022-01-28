package main

import (
	"fmt"
)

func main() {

	fruitList := []string{
		"Apple",
		"Orange",
		"Mango",
		"PineApple",
		"Guava"}

	for _, i := range fruitList {

		fmt.Println(i)

	}
}
