package main

import (
	"fmt"
	"go-projects/04.html_link_parser/parser"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	fileName := "./input/ex3.html"
	doc := parser.ParseFile(fileName)

	var linkNodes []html.Node
	parser.ParseLinkNodes(doc, &linkNodes)

	links := buildLinks(linkNodes)
	fmt.Println(links)
}

func buildLinks(linkNodes []html.Node) []Link {
	links := make([]Link, len(linkNodes))
	for i, node := range linkNodes {
		href := parser.GetHref(&node)
		text := parser.GetText(&node)
		links[i] = Link{href, text}
	}
	return links
}
