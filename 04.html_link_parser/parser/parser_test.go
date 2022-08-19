package parser

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

const htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
    <p>more content <a href="/details">here</a></p>
	<a href="/page">A link to a page</a>
</body>
</html>`

func TestGetTextShouldReturnTextContentInHtml(t *testing.T) {
	// Arrange
	doc, _ := html.Parse(strings.NewReader(htm))
	expected := "body content\n    more content here\n\tA link to a page"

	// Act
	actual := GetText(doc)

	// Assert
	if actual != expected {
		t.Errorf("Expected result to be %s, but got %s", expected, actual)
	}
}

func TestGetHrefShouldGetHrefAttrFromNode(t *testing.T) {
	// Arrange
	node := html.Node{Attr: []html.Attribute{
		{Key: "target", Val: "_blank"},
		{Key: "href", Val: "/page"},
		{Key: "referrerpolicy", Val: "same-origin"},
	}}
	expected := "/page"

	// Act
	actual := GetHref(&node)

	// Assert
	if actual != expected {
		t.Errorf("Expected result to be %s, but got %s", expected, actual)
	}
}

func TestParseLinkNodesShouldGetAllLinksInHtml(t *testing.T) {
	// Arrange
	doc, _ := html.Parse(strings.NewReader(htm))
	var linkNodes []html.Node

	// Act
	ParseLinkNodes(doc, &linkNodes)

	// Assert
	if len(linkNodes) != 2 {
		t.Errorf("Expected to return %v nodes, but got %v", 1, len(linkNodes))
	}

	for _, link := range linkNodes {
		if link.Type != html.ElementNode || link.Data != "a" {
			t.Errorf("Expected to return type %s nodes, but got %s", "a", linkNodes[0].Data)
		}
	}
}
