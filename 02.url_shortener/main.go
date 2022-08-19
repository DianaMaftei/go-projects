package main

import (
	"flag"
	"fmt"
	"go-projects/02.url_shortener/db"
	"go-projects/02.url_shortener/io"
	"go-projects/02.url_shortener/urlshort"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlF := flag.String("yaml", "paths.yaml", "Provide a yaml in the form of a list of path and url")
	jsonF := flag.String("json", "paths.json", "Provide a json in the form of a list of path and url")
	flag.Parse()

	yaml := io.ReadFile(*yamlF)
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	json := io.ReadFile(*jsonF)
	jsonHandler, err := urlshort.JSONHandler(json, yamlHandler)
	if err != nil {
		panic(err)
	}

	db := db.GetPathsFromDb()
	dbHandler := urlshort.MapHandler(db, jsonHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
