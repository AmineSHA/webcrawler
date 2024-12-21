package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {

	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("network error: %v", err)

	} else if resp.StatusCode >= 400 {
		return "", fmt.Errorf("non OK status Code: %d", resp.StatusCode)

	} else if contentType := resp.Header.Get("content-type"); !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("non HTML page:[ %s ]", contentType)

	}

	defer resp.Body.Close()

	htmlCode, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading HTML code: %v", err)
	}

	return string(htmlCode), nil

}
