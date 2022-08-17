package main

import (
	"os"
	"text/template"
)

type book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	books := []book{
		{"American Psycho", "Bret Easton Ellis", 1991},
		{"The Hound of the Baskervilles", "Arthur Conan Doyle", 1902},
		{"Camino Island", "John Grisham", 2017},
	}
	tmpl, err := template.New("test").Parse("{{range .}}{{.Title}} was written in {{.Year}} by {{.Author}} \n{{end}}")
	handleErr(err)

	err = tmpl.Execute(os.Stdout, books)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
