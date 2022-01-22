package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name   = "Jason"
	course = "Go fundamentals"
	module = "4"
	clip   = 2
)

func main() {

	fmt.Println("Name and course are set to", name, "and", course)
	fmt.Println("Module and clip are set to", module, "and", clip)
	//Giving Type Output
	fmt.Println("Type of module", reflect.TypeOf(module))
	fmt.Println("Type of name", reflect.TypeOf(name))
	//total := module + clip
	//fmt.Println("Total:-", total)

	iModule, err := strconv.Atoi(module)
	if err == nil {

		total := iModule + clip
		fmt.Println("Total:-", total)

	}

}
