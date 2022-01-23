package main

import (
	"fmt"
	"os"
)

func main() {

	//fmt.Println(os.Environ())

	//To get all the env variables one by one

	for _, env := range os.Environ() {

		fmt.Println(env)
	}
}
