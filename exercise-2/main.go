package main

import (
	"fmt"
	"net/http"
)

// mapHandler will return an http.HandlerFunc that will attempt to map any
// paths (keys in the map) to their corresponding URL (values in the map).
// If the path is not provided in the map, then the fallback http.Handler
// will be called instead.
func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// yamlHandler will parse the provided YAML and then return an http.HandlerFunc
// that will attempt to map any paths to their corresponding URL. If the
// path is not provided in the YAML, then the fallback http.Handler will be called
// instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned are related to having invalid YAML data.
func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world")
}

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := mapHandler(pathsToUrls, mux)

	//	yaml := `
	//- path: /urlshort
	//url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	//url: https://github.com/gophercises/urlshort/tree/solution
	//	`
	//	yamlHandler, err := yamlHandler([]byte(yaml), mapHandler)
	//	if err != nil {
	//		panic(err)
	//	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}
