package main

import (
	"fmt"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf(" REPORT for %s \n", baseURL)
	fmt.Println("=============================")

	sorted := sortPages(pages)
	for _, site := range sorted {
		fmt.Printf("Found %d internal links to %s times. \n", site.visits, site.page)
	}
	fmt.Printf("You visited %d pages. \n", len(pages))

}
