package dom

import (
	"strings"

	"golang.org/x/net/html"
)

// GetNodeID returns id attr of node
func GetNodeID(n *html.Node) string {
	var id string

	for _, a := range n.Attr {
		if a.Key == "id" {
			id = a.Val
			break
		}
	}

	return id
}

// ExtractNodeClasslist ...
func GetNodeClasslist(n *html.Node) []string {
	classnames := ""

	for _, a := range n.Attr {
		if a.Key != "class" {
			continue
		}
		classnames = a.Val
		break
	}

	// Split the classnames into a slice
	return strings.Fields(classnames)
}

// FindNodeByType returns first node with provided type
func FindNodeByType(n *html.Node, t string) *html.Node {
	if n == nil {
		return nil
	}

	if n.Type == html.ElementNode && n.Data == t {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found := FindNodeByType(c, t)

		if found != nil {
			return found
		}
	}

	return nil
}

// NodeContainsClasses accepts an html.Node  and a string of classnames
// and returns a bool determining if the node classlist attribute contains
// all the classes
func NodeContainsClasses(n *html.Node, c string) bool {
	classList := ExtractNodeClasslist(n)

	if arrayContains(classList, c) == true {
		return true
	}

	return false
}
