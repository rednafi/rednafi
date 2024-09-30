package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const httpResponse = `
<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:content="http://purl.org/rss/1.0/modules/content/">
	<channel>
		<title>Redowan&#39;s Reflections</title>
		<link>https://rednafi.com/</link>
		<description>Recent content on Redowan&#39;s Reflections</description>
		<image>
			<title>Redowan&#39;s Reflections</title>
			<url>https://rednafi.com/images/home/cover.webp</url>
			<link>https://rednafi.com/images/home/cover.webp</link>
		</image>
		<generator>Hugo -- gohugo.io</generator>
		<language>en</language>
		<copyright> &lt;a href=&#39;https://rednafi.com/blogroll&#39;&gt;blogroll&lt;/a&gt; • &lt;a href=&#39;https://rednafi.com/reads&#39;&gt;reads&lt;/a&gt; • &lt;a href=&#39;https://rednafi.com/uses&#39;&gt;uses&lt;/a&gt; </copyright>
		<lastBuildDate>Mon, 08 Jan 2024 00:00:00 +0000</lastBuildDate><atom:link href="https://rednafi.com/index.xml" rel="self" type="application/rss+xml" />
		<item>
			<title>Annotating args and kwargs in Python</title>
			<link>https://rednafi.com/python/annotate_args_and_kwargs/</link>
			<pubDate>Mon, 08 Jan 2024 00:00:00 +0000</pubDate>

			<guid>https://rednafi.com/python/annotate_args_and_kwargs/</guid>
			<description></description>
		</item>

		<item>
			<title>Rate limiting via Nginx</title>
			<link>https://rednafi.com/go/rate_limiting_via_nginx/</link>
			<pubDate>Sat, 06 Jan 2024 00:00:00 +0000</pubDate>

			<guid>https://rednafi.com/go/rate_limiting_via_nginx/</guid>
			<description></description>
		</item>

		<item>
			<title>Statically enforcing frozen data classes in Python</title>
			<link>https://rednafi.com/python/statically_enforcing_frozen_dataclasses/</link>
			<pubDate>Thu, 04 Jan 2024 00:00:00 +0000</pubDate>

			<guid>https://rednafi.com/python/statically_enforcing_frozen_dataclasses/</guid>
			<description></description>
		</item>

		<item>
			<title>Planning palooza</title>
			<link>https://rednafi.com/zephyr/planning_palooza/</link>
			<pubDate>Mon, 01 Jan 2024 00:00:00 +0000</pubDate>

			<guid>https://rednafi.com/zephyr/planning_palooza/</guid>
			<description></description>
		</item>

		<item>
			<title>Reminiscing CGI scripts</title>
			<link>https://rednafi.com/go/reminiscing_cgi_scripts/</link>
			<pubDate>Mon, 25 Dec 2023 00:00:00 +0000</pubDate>

			<guid>https://rednafi.com/go/reminiscing_cgi_scripts/</guid>
			<description></description>
		</item>
	</channel>
</rss>
`

func TestFetchRSS(t *testing.T) {
	tests := []struct {
		name          string
		statusCode    int
		contentType   string
		responseBody  string
		expectError   bool
		expectedError string
	}{
		{
			name:          "Valid RSS response",
			statusCode:    http.StatusOK,
			contentType:   "application/xml",
			responseBody:  httpResponse,
			expectError:   false,
			expectedError: "",
		},
		{
			name:          "Invalid content type",
			statusCode:    http.StatusOK,
			contentType:   "text/html",
			responseBody:  "<html><body>Not XML</body></html>",
			expectError:   true,
			expectedError: "Invalid content type",
		},
		{
			name:          "HTTP status error",
			statusCode:    http.StatusInternalServerError,
			contentType:   "application/xml",
			responseBody:  "",
			expectError:   true,
			expectedError: "HTTP request failed with status: 500 Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server to mock the HTTP response
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tt.contentType)
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			// Call the fetchRSS function with the test server URL
			data, err := fetchRSS(server.URL)
			if (err != nil) != tt.expectError {
				t.Fatalf("fetchRSS() error = %v, expectError %v", err, tt.expectError)
			}

			if err != nil && !strings.Contains(err.Error(), tt.expectedError) {
				t.Errorf("fetchRSS() error = %v, expected error = %s", err, tt.expectedError)
			}

			if !tt.expectError && string(data) != tt.responseBody {
				t.Errorf("fetchRSS() returned unexpected data. Got: %s, Want: %s", string(data), tt.responseBody)
			}
		})
	}
}

func TestParseRSS(t *testing.T) {
	// Call the parseRSS function with the test server URL
	rss, err := parseRSS([]byte(httpResponse))
	if err != nil {
		t.Errorf("parseRSS returned an error: %v", err)
	}

	// Check if the parsed RSS contains the expected number of items
	expected := 5
	if len(rss.Items) != expected {
		t.Errorf("parseRSS returned unexpected number of items. Got: %d, Want: %d", len(rss.Items), expected)
	}

	// Check if the first item matches the expected value
	expectedTitle := "Annotating args and kwargs in Python"
	if rss.Items[0].Title != expectedTitle {
		t.Errorf("parseRSS returned unexpected title. Got: %s, Want: %s", rss.Items[0].Title, expectedTitle)
	}
}

func TestBuildMarkdown(t *testing.T) {
	// Create a test RSS struct
	rss := RSS{
		Items: []Item{
			{
				Title:   "Annotating args and kwargs in Python",
				Link:    "https://rednafi.com/python/annotate_args_and_kwargs/",
				PubDate: "Mon, 08 Jan 2024 00:00:00 +0000",
			},
			{
				Title:   "Rate limiting via Nginx",
				Link:    "https://rednafi.com/go/rate_limiting_via_nginx/",
				PubDate: "Sat, 06 Jan 2024 00:00:00 +0000",
			},
			{
				Title:   "Statically enforcing frozen data classes in Python",
				Link:    "https://rednafi.com/python/statically_enforcing_frozen_dataclasses/",
				PubDate: "Thu, 04 Jan 2024 00:00:00 +0000",
			},
			{
				Title:   "Planning palooza",
				Link:    "https://rednafi.com/zephyr/planning_palooza/",
				PubDate: "Mon, 01 Jan 2024 00:00:00 +0000",
			},
			{
				Title:   "Reminiscing CGI scripts",
				Link:    "https://rednafi.com/go/reminiscing_cgi_scripts/",
				PubDate: "Mon, 25 Dec 2023 00:00:00 +0000",
			},
		},
	}

	// Call the buildMarkdown function with the test RSS struct
	markdown := buildMarkdown(rss, "Test title")
	expected := `Test title<div align="center">

#### Recent articles

| Title | Published On |
| ----- | ------------ |
| [Annotating args and kwargs in Python](https://rednafi.com/python/annotate_args_and_kwargs/) | Mon, 08 Jan 2024 |
| [Rate limiting via Nginx](https://rednafi.com/go/rate_limiting_via_nginx/) | Sat, 06 Jan 2024 |
| [Statically enforcing frozen data classes in Python](https://rednafi.com/python/statically_enforcing_frozen_dataclasses/) | Thu, 04 Jan 2024 |
| [Planning palooza](https://rednafi.com/zephyr/planning_palooza/) | Mon, 01 Jan 2024 |
| [Reminiscing CGI scripts](https://rednafi.com/go/reminiscing_cgi_scripts/) | Mon, 25 Dec 2023 |
</div>`

	// Check if the generated markdown matches the expected value
	if markdown != expected {
		t.Errorf("buildMarkdown returned unexpected markdown. Got: %s, Want: %s", markdown, expected)
	}

}
