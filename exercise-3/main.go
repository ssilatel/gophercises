package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Story map[string]Page

type Page struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	// Decoding the JSON file
	file, err := os.Open("story.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(file)
	var story Story
	err = dec.Decode(&story)
	if err != nil {
		panic(err)
	}

	http.Handle("/", new(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
