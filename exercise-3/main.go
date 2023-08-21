package main

import (
	"encoding/json"
	"html/template"
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
	t := `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
		<style>
			.main-div {
				display: flex;
				flex-direction: column;
				border: 5px solid;
				margin: auto;
				width: 600px;
				padding: 10px;
				text-align: center;
				align-items: center;
			}

			.options {
				width: 400px;
			}

			ul {
				border-top: 3px dotted #ddd;
			}
		</style>
	</head>
	<body>
		<div class="main-div">
			<h1>{{.Title}}</h1>
			<div>
				{{range .Story}}
					<p>{{.}}</p>
				{{end}}
			</div>
			{{if .Options}}
			<div class="options">
				<ul>
					{{range .Options}}
						<li><a href="/{{.Arc}}">{{.Text}}</a></li>
					{{end}}
				</ul>
			</div>
			{{end}}
		</div>
	</body>
</html>
`

	tmpl := template.Must(template.New("").Parse(t))
	tmpl.Execute(w, h.s[r.URL.Path[1:]])
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

	h := handler{s: story}
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
