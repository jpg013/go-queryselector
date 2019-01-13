package queryselector

/*
type nodeFn func(t *html.Node)

func classTreeWalker(n *html.Node, c string, s []*html.Node) []*html.Node {
	if n == nil {
		return s
	}

	if NodeContainsClasses(n, c) == true {
		s = append(s, n)
	}

	if n.FirstChild != nil {
		s = classTreeWalker(n.FirstChild, c, s)
	}

	s = classTreeWalker(n.NextSibling, c, s)

	return s
}

// QuerySelectorAll function
func QuerySelectorAll(n *html.Node, c string) []*html.Node {
	return classTreeWalker(n, c, make([]*html.Node, 0))
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
*/
