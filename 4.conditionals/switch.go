package main

import (
	"fmt"
)

func main() {

	switch 1 {

	case 1:
		fmt.Println("Value is 1")
		fallthrough
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")

	default:
		fmt.Println("Out of range")

	}
}
