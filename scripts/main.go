package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	rssURL          = "https://rednafi.com/index.xml"
	outputFile      = "README.md"
	dateFormatLimit = 16 // Limits the date format to the first 16 characters
	header          = `<div align="center">
Wandering dilettante with a flair for 1s and 0s <br>
Find my musings at <a href="https://rednafi.com/" rel="me">rednafi.com</a>
<div>`
)

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Items []Item `xml:"item"`
	} `xml:"channel"`
}

func fetchRSS(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func parseRSS(data []byte) (RSS, error) {
	var rss RSS
	err := xml.Unmarshal(data, &rss)
	return rss, err
}

func buildMarkdown(rss RSS, header string) string {
	markdown := fmt.Sprintf("%s\n\n#### Recent articles\n\n", header)
	markdown += "| Title | Published On |\n| ----- | ------------ |\n"

	for _, item := range rss.Channel.Items[:5] {
		markdown += fmt.Sprintf(
			"| [%s](%s) | %s |\n", item.Title, item.Link, item.PubDate[:dateFormatLimit],
		)
	}
	return markdown
}

func writeToFile(content string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func main() {
	rssData, err := fetchRSS(rssURL)
	if err != nil {
		fmt.Printf("Error fetching RSS: %s\n", err)
		return
	}

	rss, err := parseRSS(rssData)
	if err != nil {
		fmt.Printf("Error parsing RSS: %s\n", err)
		return
	}

	markdown := buildMarkdown(rss, header)
	if err := writeToFile(markdown, outputFile); err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	log.Printf("Successfully written to %s\n\n", outputFile)

	fmt.Println("Markdown content ")
	fmt.Printf("================\n\n")
	fmt.Println(markdown)
}
