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

}
