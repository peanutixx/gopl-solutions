// Exercise 5.4:
// Extend the visit function so that it extracts other kinds of links from the document,
// such as images, scripts, and style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text: %v\n", err)
		os.Exit(1)
	}

	for _, link := range extendedVisit(nil, doc) {
		fmt.Println(link)
	}
}

func extendedVisit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		if node.Data == "a" || node.Data == "link" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		if node.Data == "img" || node.Data == "script" || node.Data == "style" {
			for _, a := range node.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = extendedVisit(links, c)
	}
	return links
}
