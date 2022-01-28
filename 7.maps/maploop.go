package main

import (
	"fmt"
)

func main() {

	testMap := map[string]int{

		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"J": 7,
	}
	fmt.Println(testMap)

	for mapKey, mapVal := range testMap {

		fmt.Println("Key is: Value is: \n", mapKey, mapVal)
	}
}
