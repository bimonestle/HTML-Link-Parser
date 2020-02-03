package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represent a link (<a href="...") in an HTML document.
type Link struct {
	Href string
	Text string
}

// 1. Create something to read an HTML Document or
// parse an HTML Document.

// Parse will take in an HTML document and will return
// a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

// Depth First Search algorithm.
// Will return every single nodes of HTML Document
func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)

	// This loop will check every single nodes
	// from root element until it hits the last element
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
