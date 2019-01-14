package queryselector

import (
	"golang.org/x/net/html"
)

type nodeFn func(t *html.Node)

// GetElementByClassname accepts an html.Node and classname string and
// returns the first html node that contains the classname.
func GetElementByClassname(doc *html.Node, c string) *html.Node {
	fn := func(n *html.Node) bool {
		return ElementContainsClass(n, c)
	}

	return WalkTheDOM(doc, fn)
}

// GetElementsByClassname accepts an html.Node and classname string and
// returns a slice of all html nodes that contains the classname.
func GetElementsByClassname(doc *html.Node, c string) []*html.Node {
	fn := func(n *html.Node) bool {
		return ElementContainsClass(n, c)
	}

	return WalkTheDOMAll(doc, fn, make([]*html.Node, 0))
}
