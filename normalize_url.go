package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("could not parse url")
	}

	normalizedURL := strings.Trim(strings.ToLower(parsedURL.Host+parsedURL.Path), "/")

	return normalizedURL, nil

}
