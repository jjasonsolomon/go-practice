package main

import (
	"fmt"
)

func main() {

	biggestNumber := bigNumberFinder(1, 4, 3, 10, 8, 6)
	fmt.Println(biggestNumber)

}

func bigNumberFinder(bigNo ...int) int {

	number := bigNo[0]
	for _, i := range bigNo {

		if i > number {
			number = i
		}

	}

	return number

}
