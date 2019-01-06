package queryselector

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

var b, ioErr = ioutil.ReadFile("./test.html") // just pass the file name

var doc, readerErr = html.Parse(bytes.NewReader(b))

var testNode = doc.FirstChild.FirstChild.NextSibling.FirstChild

func TestParseNodeID(t *testing.T) {
	id := ParseNodeID(testNode)

	if id != "testid_123" {
		t.Errorf("Node ID is incorrect, got: %s, wanted: %s.", id, "testid_123")
	}
}

func TestParseNodeClasslist(t *testing.T) {
	cxs := ParseNodeClasslist(testNode)

	if cxs[0] != "testclass1" {
		t.Errorf("Node Classlist is incorrect, got: %s, wanted: %s.", cxs[0], "testclass1")
	}

	if cxs[1] != "testclass2" {
		t.Errorf("Node Classlist is incorrect, got: %s, wanted: %s.", cxs[0], "testclass2")
	}
}
