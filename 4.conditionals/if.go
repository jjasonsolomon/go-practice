package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 20

	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("b is greater a")
	}

}
