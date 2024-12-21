package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	defer cfg.wg.Done()

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
			<-cfg.concurrencyControl
			return
		}

		fmt.Printf("Current page: %s \n", rawCurrentURL)
		firstVisit := cfg.addPageVisit(normalizedCurrent)

		if !firstVisit {
			<-cfg.concurrencyControl
			return
		}

		htmlPage, err := getHTML(rawCurrentURL)
		if err != nil && strings.Contains(err.Error(), "non HTML page:") {
			<-cfg.concurrencyControl
			return
		}

		urlList, err := getURLsFromHTML(htmlPage, cfg.baseURL.String())
		if err != nil {
			<-cfg.concurrencyControl
			return
		}

		for _, url := range urlList {
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
			cfg.concurrencyControl <- struct{}{}
		}

	}
	<-cfg.concurrencyControl
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	isFirst = true

	if _, ok := cfg.pages[normalizedURL]; !ok {
		cfg.pages[normalizedURL] = 1
		return isFirst
	}
	cfg.pages[normalizedURL]++
	return !isFirst

}
