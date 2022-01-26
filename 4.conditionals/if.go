package main

import (
	"fmt"
)

func main() {

	//a := 10
	//b := 20

	//Simple Initialisation
	if a, b := 10, 20; a > b {
		fmt.Println("a is greater than b")
		
	} else if a == b {
		fmt.Println("a and b are equal")
	} else {
		
		fmt.Println("b is greater a")
	}

}
