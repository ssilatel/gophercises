package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	//	file, err := os.Open("ex1.html")
	//	if err != nil {
	//		panic(err)
	//	}

	htmlFile, err := ioutil.ReadFile("ex1.html")
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(strings.NewReader(string(htmlFile)))
	if err != nil {
		panic(err)
	}

	var parseHTML func(*html.Node)
	parseHTML = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseHTML(c)
		}
	}
	parseHTML(doc)
}
