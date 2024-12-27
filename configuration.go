package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
	counterPages       int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	isFirst = true

	if _, ok := cfg.pages[normalizedURL]; !ok {
		cfg.pages[normalizedURL] = 1
		cfg.counterPages++
		return isFirst
	}
	cfg.pages[normalizedURL]++
	return !isFirst

}

func configSetup(rawBaseURL string, maxConcurrency, totalPages int) (*config, error) {
	parsedBaseURL, err := url.Parse(rawBaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing baseUrl %s", rawBaseURL)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           totalPages,
		counterPages:       0,
	}, nil
}
