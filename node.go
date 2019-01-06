package queryselector

import (
	"strings"

	"golang.org/x/net/html"
)

// ParseNodeID returns id attr of node
func ParseNodeID(n *html.Node) string {
	var id string

	for _, a := range n.Attr {
		if a.Key == "id" {
			id = a.Val
			break
		}
	}

	return id
}

// GetNodeAttribute accepts an html.Node pointer and a
// key string and returns the attribute value
func GetNodeAttribute(n *html.Node, k string) string {
	var s string

	for _, a := range n.Attr {
		if a.Key == k {
			s = a.Val
			break
		}
	}

	return s
}

// ParseNodeClasslist accepts an html.Node pointer and returns
// a the nodes class attributes as a slice of strings
func ParseNodeClasslist(n *html.Node) []string {
	classnames := GetNodeAttribute(n, "class")

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
	classList := ParseNodeClasslist(n)

	if arrayContains(classList, c) == true {
		return true
	}

	return false
}

func arrayContains(cx []string, s string) bool {
	matchCount := 0

	split := strings.Fields(s)

	if len(split) == 0 {
		return false
	}

	for _, x := range split {
		match := false

		for _, c := range cx {
			if x == c {
				match = true
				break
			}
		}

		if match == true {
			matchCount++
		}
	}

	if matchCount == len(split) {
		return true
	}

	return false
}
