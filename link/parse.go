package link

import "io"

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
	return nil, nil
}
