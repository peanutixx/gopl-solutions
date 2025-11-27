// Exercise 5.2:
// Write a function to populate a mapping from element names -- p, div, span,
// and so on -- to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "polulate: %v\n", err)
		os.Exit(1)
	}
	result := make(map[string]int)
	populate(result, doc)

	for name, count := range result {
		fmt.Printf("%s: %d\n", name, count)
	}
}

func populate(result map[string]int, node *html.Node) {
	if node.Type == html.ElementNode {
		result[node.Data]++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		populate(result, c)
	}
}
