package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
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

		fmt.Printf("starting crawl of: %v \n", os.Args[1])
		parsedBaseURL, err := url.Parse(os.Args[1])
		if err != nil {
			fmt.Printf("Error parsing baseUrl")
			os.Exit(1)
		}
		var cfg config = config{
			pages:              make(map[string]int),
			baseURL:            parsedBaseURL,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, 5),
			wg:                 &sync.WaitGroup{},
		}

		cfg.crawlPage(os.Args[1])
		cfg.wg.Wait()
		fmt.Print("----------------------- \n")
		for k, v := range cfg.pages {
			fmt.Printf("You visited %s %d times. \n", k, v)
		}

	}

}
