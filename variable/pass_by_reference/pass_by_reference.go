package main

import (
	"fmt"
)

func main() {

	name := "Jason"
	course := "Go"
	fmt.Println("\nHi", name, "is currently learning", course, "course")
	updateCourse(&course)
	fmt.Println("\nHi", name, "is currently learning", course, "course")

}

func updateCourse(course *string) string {
	*course = "Docker"
	fmt.Println("\nUpdated the course to", *course)
	return *course

}
