package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

func ParseFile(fileName string) *html.Node {
	f, err := os.Open(fileName)
	handleError(err, fmt.Sprintf("Unable to read file %s", fileName))

	doc, err := html.Parse(f)
	handleError(err, fmt.Sprintf("Unable to parse file %s", fileName))
	return doc
}

func ParseLinkNodes(node *html.Node, nodes *[]html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		*nodes = append(*nodes, *node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ParseLinkNodes(child, nodes)
	}
}

func GetHref(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func GetText(node *html.Node) string {
	var sb strings.Builder
	buildText(node, &sb)
	return strings.TrimSpace(sb.String())
}

func buildText(node *html.Node, sb *strings.Builder) {
	if node.Type == html.TextNode {
		sb.WriteString(node.Data)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		buildText(child, sb)
	}
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
		os.Exit(1)
	}
}
