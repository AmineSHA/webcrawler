package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLSFromHTML(t *testing.T) {

	tests := []struct {
		name          string
		htmlBody      string
		rawBaseUrl    string
		expected      []string
		errorContains string
	}{
		{
			name:       "absolute URL",
			rawBaseUrl: "https://blog.boot.dev",
			htmlBody: `
<html>
	<body>
		<a href="https://blog.boot.dev">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:       "relative URL",
			rawBaseUrl: "https://blog.boot.dev",
			htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:       "absolute and relative URLs",
			rawBaseUrl: "https://blog.boot.dev",
			htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:       "no href",
			rawBaseUrl: "https://blog.boot.dev",
			htmlBody: `
<html>
	<body>
		<a>
			<span>Boot.dev></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:       "bad HTML",
			rawBaseUrl: "https://blog.boot.dev",
			htmlBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:       "invalid href URL",
			rawBaseUrl: "https://blog.boot.dev",
			htmlBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:       "handle invalid base URL",
			rawBaseUrl: `:\\invalidBaseURL`,
			htmlBody: `
<html>
	<body>
		<a href="/path">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected:      nil,
			errorContains: "couldn't parse base URL",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.htmlBody, tc.rawBaseUrl)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return

			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}

		})
	}

}
