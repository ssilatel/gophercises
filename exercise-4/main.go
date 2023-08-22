package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func (l *Link) printLink() {
	fmt.Println("href: " + l.Href)
	fmt.Println("text: " + l.Text)
	fmt.Println()
}

func main() {
	filename := flag.String("f", "ex1.html", "HTML file to parse")
	flag.Parse()

	htmlFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(strings.NewReader(string(htmlFile)))
	if err != nil {
		panic(err)
	}

	links := []Link{}

	var parseHtml func(*html.Node)
	parseHtml = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, Link{Href: a.Val, Text: n.FirstChild.Data})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseHtml(c)
		}
	}
	parseHtml(doc)

	for _, l := range links {
		l.printLink()
	}
}
