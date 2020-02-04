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

	// A slice of Link struct or object
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))

		// fmt.Println(node)
	}
	// This dfs() below will return every single nodes of an
	// HTML Document
	// dfs(doc, "")
	return links, nil
}

// Get the href value from an HTML Document
// and store it to the Link struct {Href string, Text string}
func buildLink(n *html.Node) Link {
	//
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

// Get text value inside an <a> from
// an HTML Document. Pass it to buildLink()
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	// Otherwise, the type given is for example
	// a comment or Doctype. Do this
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return ret
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
