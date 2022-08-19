package main

import (
	"encoding/json"
	"fmt"
	"go-projects/03.choose_your_own_adventure/handler"
	"go-projects/03.choose_your_own_adventure/io"
	"go-projects/03.choose_your_own_adventure/model"
	"net/http"
)

func main() {
	mux := defaultMux()

	adventure := parseJson()

	htmlHandler := handler.HTMLHandler(adventure, mux)

	http.ListenAndServe(":8080", htmlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "/intro", http.StatusFound)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintln(w, "404 Page Not Found")
	}
}

func parseJson() map[string]model.Page {
	file := io.ReadFile("gopher.json")

	parsedFile := make(map[string]model.Page)

	err := json.Unmarshal(file, &parsedFile)
	if err != nil {
		return nil
	}
	return parsedFile
}
