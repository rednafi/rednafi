package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	rssURL          = "https://rednafi.com/index.xml"
	outputFile      = "README.md"
	dateFormatLimit = 16
	header          = `<div align="center">
Roving amateur with a flair for words and wires. <br>
Find my musings at <a href="https://rednafi.com/" rel="me">rednafi.com</a>
</div>`
)

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

type RSS struct {
	Items []Item `xml:"channel>item"`
}

func fetchRSS(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/xml") &&
		!strings.Contains(contentType, "text/xml") {
		return nil, fmt.Errorf("Invalid content type: %s", contentType)
	}

	return body, nil
}

func parseRSS(data []byte) (RSS, error) {
	var rss RSS
	err := xml.Unmarshal(data, &rss)
	return rss, err
}

func buildMarkdown(rss RSS, header string) string {
	markdown := fmt.Sprintf("%s", header)
	markdown += `<div align="center">`
	markdown += "\n\n#### Recent articles\n\n"
	markdown += "| Title | Published On |\n| ----- | ------------ |\n"

	for _, item := range rss.Items[:5] {
		markdown += fmt.Sprintf("| [%s](%s) | %s |\n", item.Title, item.Link, item.PubDate[:dateFormatLimit])
	}

	markdown += `</div>`
	return markdown
}

func writeToFile(content, filename string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func main() {
	rssData, err := fetchRSS(rssURL)
	if err != nil {
		log.Fatalf("Error fetching RSS: %v\n", err)
	}

	rss, err := parseRSS(rssData)
	if err != nil {
		log.Fatalf("Error parsing RSS: %v\n", err)
	}

	markdown := buildMarkdown(rss, header)
	if err := writeToFile(markdown, outputFile); err != nil {
		log.Fatalf("Error writing to file: %v\n", err)
	}

	log.Printf("Successfully written to %s\n", outputFile)

	fmt.Println("Markdown content:")
	fmt.Println("================")
	fmt.Println(markdown)
}
