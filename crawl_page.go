package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	parsedBaseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing baseUrl")
		return
	}
	parsedCurrentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing baseUrl")
		return
	}

	if parsedBaseUrl.Host == parsedCurrentUrl.Host {

		normalizedCurrent, err := normalizeURL(rawCurrentURL)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("Current page: %s \n", rawCurrentURL)

		if _, ok := pages[normalizedCurrent]; !ok {
			pages[normalizedCurrent] = 1
		} else {
			pages[normalizedCurrent]++
			return
		}

		htmlPage, err := getHTML(rawCurrentURL)
		if err != nil && strings.Contains(err.Error(), "non HTML page:") {
			return
		}

		urlList, err := getURLsFromHTML(htmlPage, rawBaseURL)
		if err != nil {
			return
		}

		for _, url := range urlList {
			crawlPage(rawBaseURL, url, pages)
		}

	}

}
