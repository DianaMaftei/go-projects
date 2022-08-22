package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) []Link {
	doc, err := html.Parse(r)
	handleError(err, fmt.Sprintf("Unable to parse html"))

	var linkNodes []html.Node
	parseLinkNodes(doc, &linkNodes)

	return buildLinks(linkNodes)
}

func parseLinkNodes(node *html.Node, nodes *[]html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		*nodes = append(*nodes, *node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		parseLinkNodes(child, nodes)
	}
}

func getHref(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func getText(node *html.Node) string {
	var sb strings.Builder
	buildText(node, &sb)
	return strings.TrimSpace(sb.String())
}

func buildLinks(linkNodes []html.Node) []Link {
	links := make([]Link, len(linkNodes))
	for i, node := range linkNodes {
		href := getHref(&node)
		if len(href) == 0 {
			continue
		}
		text := getText(&node)
		links[i] = Link{href, text}
	}
	return links
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
