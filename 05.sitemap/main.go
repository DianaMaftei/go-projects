package main

import (
	"encoding/xml"
	"go-projects/04.html_link_parser/parser"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const rootUrl = "https://www.calhoun.io"
const domain = "calhoun.io"

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func main() {
	siteMapLinks := make(map[string]bool)
	siteMapLinks[rootUrl] = false
	hasUnvisitedPages := true

	for {
		getLinksForAllUrls(siteMapLinks)
		hasUnvisitedPages = false
		for _, val := range siteMapLinks {
			if !val {
				hasUnvisitedPages = true
			}
		}
		if !hasUnvisitedPages {
			break
		}
	}

	urls := buildUrlList(siteMapLinks)

	composeSiteMap("sitemap.xml", urls)
}

func buildUrlList(siteMapLinks map[string]bool) []Url {
	keys := getSortedUrls(siteMapLinks)

	urls := make([]Url, len(keys))

	for i, k := range keys {
		urls[i] = Url{Loc: k}
	}
	return urls
}

func getSortedUrls(siteMapLinks map[string]bool) []string {
	keys := make([]string, 0, len(siteMapLinks))

	for k := range siteMapLinks {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func composeSiteMap(filename string, urls []Url) {
	header := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
`

	out, _ := xml.MarshalIndent(urls, " ", "  ")
	body := string(out)

	footer := `
</urlset>`

	writeToFile(filename, header+body+footer)
}

func writeToFile(filename, text string) {
	textBytes := []byte(text)

	ioutil.WriteFile(filename, textBytes, 0666)
}
func getLinksForAllUrls(siteMapLinks map[string]bool) {
	for key, val := range siteMapLinks {
		if !val {
			getLinksFromPage(siteMapLinks, key)
			siteMapLinks[key] = true
		}
	}
}

func getLinksFromPage(sameDomainLinks map[string]bool, url string) {
	links := getHtmlLinks(url)

	for _, link := range links {
		url := link.Href
		if strings.HasPrefix(link.Href, "/") {
			url = rootUrl + link.Href
		}

		if isFromSameDomain(url) && !sameDomainLinks[url] {
			if _, exists := sameDomainLinks[url]; !exists {
				sameDomainLinks[url] = false
			}
		}
	}
}

func isFromSameDomain(url string) bool {
	return strings.HasPrefix(url, "https") && strings.Contains(url, domain)
}

func getHtmlLinks(page string) []parser.Link {
	resp, err := http.Get(page)

	handleErr(err)
	defer resp.Body.Close()

	links := parser.Parse(resp.Body)

	return links
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
