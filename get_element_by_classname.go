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
