package main

import (
	"fmt"
)

func main() {

	leagueTitles := make(map[string]int)

	leagueTitles["Sunderland"] = 6
	leagueTitles["NewCastle"] = 4

	recentHead2HeadWins := map[string]int{

		"Sunderland": 6,
		"NewCastle":  0,
	}

	fmt.Println("League titles: \nRecent Head to heads: \n", leagueTitles, recentHead2HeadWins)

}
