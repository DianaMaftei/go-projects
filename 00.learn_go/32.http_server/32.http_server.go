package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/hello", handleHello)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	hasName := r.URL.Query().Has("name")
	fmt.Println("Called '/'")
	if hasName {
		name := r.URL.Query().Get("name")
		io.WriteString(w, "Welcome, "+name+"!")
	} else {
		io.WriteString(w, "Welcome, You!")
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called '/hello'")
	io.WriteString(w, "Hello, world!\n")
}
