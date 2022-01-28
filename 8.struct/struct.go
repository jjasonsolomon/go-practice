package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	goFundamentals := courseMeta{

		author: "Nigel Poulton",
		level:  "Intermediate",
		rating: 3,
	}

	fmt.Println(goFundamentals)

	//To print only the author

	fmt.Println("Author of Go Fundamentals is", goFundamentals.author)

	//changing the rating

	goFundamentals.rating = 5

	fmt.Println("Current course rating is", goFundamentals.rating)

}
