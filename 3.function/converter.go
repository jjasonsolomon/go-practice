package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "jason"
	course := "go fundamentals"
	fmt.Println(convert(name, course))
}

func convert(name, course string) (str1, str2 string) {

	name = strings.ToUpper(name)
	course = strings.Title(course)
	return name, course
}
