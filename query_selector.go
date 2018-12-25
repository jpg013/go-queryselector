package dom

import (
	"golang.org/x/net/html"
)

type nodeFn func(t *html.Node)

func recurseDocument(n *html.Node, fn nodeFn) {
	if n == nil {
		return
	}

	fn(n)

	if n.FirstChild != nil {
		recurseDocument(n.FirstChild, fn)
	}

	if n.NextSibling != nil {
		recurseDocument(n.NextSibling, fn)
	}
}

// QuerySelectorAll function
func QuerySelectorAll(doc *html.Node, c string) []*html.Node {
	pred := func(n *html.Node) {

	}

	recurseDocument(doc, pred)

	var mySlice []*html.Node

	return mySlice
	/*
		matches := []*html.Node

		if n == nil {
			return matches
		}

		if NodeContainsClasses(n, c) == true {
			matches = append(matches, n)
		}

		matches = QuerySelectorAll(n.FirstChild, c, matches)
		matches = QuerySelectorAll(n.NextSibling, c, matches)

		return matches
	*/
}

// QuerySelector function
func QuerySelector(n *html.Node, c string) *html.Node {
	if n == nil {
		return nil
	}

	if NodeContainsClasses(n, c) == true {
		return n
	}

	found := QuerySelector(n.FirstChild, c)

	if found != nil {
		return found
	}

	found = QuerySelector(n.NextSibling, c)

	if found != nil {
		return found
	}

	return nil
}
