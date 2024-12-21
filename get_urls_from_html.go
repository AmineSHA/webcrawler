package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	foundURLS := make([]string, 0)
	for n := range doc.Descendants() {

		if n.Type == html.ElementNode && n.DataAtom == atom.A {

			for _, a := range n.Attr {
				if a.Key == "href" {
					foundURLS = append(foundURLS, a.Val)
					break
				}
			}
		}
	}
	var resultURLS []string

	for _, foundURL := range foundURLS {
		parsed, err := url.Parse(foundURL)

		if err != nil {
			fmt.Printf("couldn't parse href '%v': %v\n", foundURL, err)
			continue
		}
		resolvedURL := baseURL.ResolveReference(parsed)

		resultURLS = append(resultURLS, resolvedURL.String())

	}

	return resultURLS, nil

}
