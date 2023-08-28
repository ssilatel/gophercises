package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ssilatel/gophercises/link"

	"golang.org/x/net/html"
)

func parseHtml(url string) []link.Link {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(url)
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}
	links := link.FindLinks(doc)
	return links
}

func contains(listToSearch []string, str string) bool {
	for _, i := range listToSearch {
		if i == str {
			return true
		}
	}
	return false
}

func containsSubstring(str, substr string) bool {
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func cleanLinks(anchorTags []link.Link, url string) []string {
	var links []string
	for _, a := range anchorTags {
		if len(a.Href) > 0 && a.Href != "" && !containsSubstring(a.Href, "twitter") {
			if a.Href[0] == '/' {
				links = append(links, url+a.Href)
			} else if a.Href[0] == 'h' {
				links = append(links, a.Href)
			}
		}
	}
	return links
}

func visitLinks(linksToVisit []string) []string {
	var newLinks []string
	for _, l := range linksToVisit {
		anchorTags := parseHtml(l)
		links := cleanLinks(anchorTags, l)
		newLinks = append(newLinks, links...)
	}
	return newLinks
}

func main() {
	urlFlag := flag.String("u", "https://www.calhoun.io", "URL to scan")
	flag.Parse()

	anchorTags := parseHtml(*urlFlag)
	rootLinks := cleanLinks(anchorTags, *urlFlag)

	visitedLinks := visitLinks(rootLinks)

	for _, l := range visitedLinks {
		fmt.Println(l)
	}
}
