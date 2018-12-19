package main

import (
	"fmt"
	"net/http"

	"github.com/djackreuter/urlshortner"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/github":     "https://github.com/djackreuter",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshortner
  url: https://github.com/djackreuter/urlshortner
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshortner.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Testing")
}
