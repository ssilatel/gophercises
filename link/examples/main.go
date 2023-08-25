package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"

	"github.com/ssilatel/gophercises/link"
)

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

	links := link.FindLinks(doc)

	fmt.Println(links)
}
