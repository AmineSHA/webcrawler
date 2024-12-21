package main

import (
	"fmt"
	"os"
)

func main() {

	progArgs := os.Args[1:]
	nbArgs := len(progArgs)

	if nbArgs < 1 {
		fmt.Println("no website provided")
		os.Exit(1)

	} else if nbArgs > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)

	} else {
		pages := make(map[string]int)
		fmt.Printf("starting crawl of: %v \n", os.Args[1])
		crawlPage(os.Args[1], os.Args[1], pages)
		fmt.Print("----------------------- \n")
		for k, v := range pages {
			fmt.Printf("You visited %s %d times. \n", k, v)
		}

	}

}
