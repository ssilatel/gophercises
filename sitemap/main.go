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

func main() {
	urlFlag := flag.String("u", "https://www.calhoun.io", "URL to scan")
	flag.Parse()

	res, err := http.Get(*urlFlag)
	if err != nil {
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
	fmt.Println(links)
}
