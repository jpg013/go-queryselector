package queryselector

import (
	"strings"

	"golang.org/x/net/html"
)

// DomPredicate func type takes an *html.Node and returns bool
type DomPredicate func(n *html.Node) bool

// GetDocumentBody recurses *html.Node until a <body> node is found
// or returns nil
func GetDocumentBody(doc *html.Node) *html.Node {
	fn := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "body"
	}

	return WalkTheDOM(doc, fn)
}

// WalkTheDOM recurses *html.Node returning first node where
// DomPredicate passes
func WalkTheDOM(n *html.Node, fn DomPredicate) *html.Node {
	if fn(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found := WalkTheDOM(c, fn)

		if found != nil {
			return found
		}
	}

	return nil
}

// ElementContainsClass accepts an html.Node  and a string of classnames
// and returns a bool determining if the node classlist attribute contains
// all the classes
func ElementContainsClass(n *html.Node, c string) bool {
	cx := SplitNodeClasslist(n)
	fields := strings.Fields(c)
	counter := 0

	for _, x := range fields {
		if ContainsString(cx, x) {
			counter++
		}
	}

	return len(fields) == counter
}

// SplitNodeClasslist takes an *html.Node and splits
// the nodes class string and returns the string slice
func SplitNodeClasslist(n *html.Node) []string {
	classnames := GetAttribute(n, "class")

	// Split the classnames into a slice
	return strings.Fields(classnames)
}

// GetAttribute accepts an html.Node pointer and a
// key string and returns the attribute value
func GetAttribute(n *html.Node, k string) string {
	var s string

	for _, a := range n.Attr {
		if a.Key == k {
			s = a.Val
			break
		}
	}

	return s
}

// ParseElementID returns id string attr of given node
func ParseElementID(n *html.Node) string {
	var id string

	for _, a := range n.Attr {
		if a.Key == "id" {
			id = a.Val
			break
		}
	}

	return id
}
