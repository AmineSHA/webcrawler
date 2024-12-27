package main

import (
	"fmt"
	"net/url"
	"strings"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.wg.Add(1)

	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	if cfg.counterPages >= cfg.maxPages {
		return
	}

	parsedCurrentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing baseUrl")
		<-cfg.concurrencyControl
		return
	}

	if cfg.baseURL.Host == parsedCurrentUrl.Host {

		normalizedCurrent, err := normalizeURL(rawCurrentURL)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("Current page: %s \n", rawCurrentURL)
		firstVisit := cfg.addPageVisit(normalizedCurrent)

		if !firstVisit {
			return
		}

		htmlPage, err := getHTML(rawCurrentURL)
		if err != nil && strings.Contains(err.Error(), "non HTML page:") {
			return
		}

		urlList, err := getURLsFromHTML(htmlPage, cfg.baseURL.String())
		if err != nil {
			<-cfg.concurrencyControl
			return
		}

		for _, url := range urlList {

			go cfg.crawlPage(url)
			cfg.concurrencyControl <- struct{}{}
		}

	}
}
