// Exercise 5.1:
// Change the findlinks program to traverse the n.FirstChild linked list using
// recursive calls to visit instead of a loop.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// preorder traversal
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if node.FirstChild != nil {
		links = visit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		links = visit(links, node.NextSibling)
	}
	return links
}
