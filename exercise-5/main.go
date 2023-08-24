package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := flag.String("u", "https://www.calhoun.io", "URL to scan")
	flag.Parse()

	res, err := http.Get(*url)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)
}
