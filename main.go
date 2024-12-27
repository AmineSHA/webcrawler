package main

import (
	"fmt"
	"log"

	"os"
	"strconv"
)

type args struct {
	website        string
	maxConcurrency int
	maxPages       int
}

func main() {

	progArgs := os.Args[1:]
	nbArgs := len(progArgs)

	var arguments args = args{
		website:        "",
		maxConcurrency: 1,
		maxPages:       15,
	}

	if nbArgs > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)

	} else if nbArgs > 2 {
		conc, _ := strconv.Atoi(progArgs[1])
		arguments.maxConcurrency = conc
		maxP, _ := strconv.Atoi(progArgs[2])
		arguments.maxPages = maxP

	} else if nbArgs > 1 {
		conc, _ := strconv.Atoi(progArgs[1])
		arguments.maxConcurrency = conc
	} else if nbArgs < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	arguments.website = progArgs[0]

	fmt.Printf("starting crawl of: %v \n", os.Args[1])

	cfg, err := configSetup(arguments.website, arguments.maxConcurrency, arguments.maxPages)
	if err != nil {
		log.Fatalf("%v", err)
	}

	cfg.crawlPage(arguments.website)
	cfg.wg.Wait()
	printReport(cfg.pages, cfg.baseURL.String())

}
