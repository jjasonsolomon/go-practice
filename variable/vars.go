package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	name := "Jason"
	course := "Go fundamentals"
	module := "4"
	clip := 2

	fmt.Println("Name and course are set to", name, "and", course)
	fmt.Println("Module and clip are set to", module, "and", clip)
	//Giving Type Output
	fmt.Println("Type of module", reflect.TypeOf(module))
	fmt.Println("Type of name", reflect.TypeOf(name))

	iModule, err := strconv.Atoi(module)
	if err == nil {

		total := iModule + clip
		fmt.Println("Total:-", total)

	}

}
