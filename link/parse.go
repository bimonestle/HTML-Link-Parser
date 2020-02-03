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

	// 1. Find <a> nodes in an HTML Document
	// 2. foreach link node
	// 	2a. Store value to a Link struct
	// 3. return the Links
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	// This dfs() will return every single nodes of an
	// HTML Document
	// dfs(doc, "")
	return nil, nil
}

// Get every single <a> node / element
// from an HTML Document
func linkNodes(n *html.Node) []*html.Node {

	// This section checks:
	// if the node is an <a> already, return it.
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	// Otherwise it has to go through the
	// whole document / every single node and check
	// if it's a link
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
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
