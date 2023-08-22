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

var links []Link

func (l *Link) printLink() {
	fmt.Println("href: " + l.Href)
	fmt.Println("text: " + l.Text)
	fmt.Println()
}

func parseText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	} else {
		return ""
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

func parseHtml(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var text string
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text += parseText(c)
		}
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, Link{Href: a.Val, Text: text})
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHtml(c)
	}
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
	parseHtml(doc)

	for _, l := range links {
		l.printLink()
	}
}
