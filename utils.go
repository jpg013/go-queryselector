package queryselector

// ContainsString returns true if string is
// an element of []string, else false
func ContainsString(cx []string, s string) bool {
	for _, x := range cx {
		if x == s {
			return true
		}
	}
	return false
}
