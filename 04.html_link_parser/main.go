package main

import (
	"fmt"
	"go-projects/04.html_link_parser/parser"
	"os"
)

func main() {
	fileName := "./input/ex3.html"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	links := parser.Parse(f)

	fmt.Println(links)
}
