package queryselector

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

var b, ioErr = ioutil.ReadFile("./test_doc.html")

var doc, readerErr = html.Parse(bytes.NewReader(b))

func TestGetElementByClassname(t *testing.T) {
	n := GetElementByClassname(doc, "header-links")

	if n == nil {
		t.Errorf("GetElementByClassname returned incorrect nil")
	}
}

func TestGetElementByClassnameMultiple(t *testing.T) {
	el := GetElementByClassname(doc, "header-links primary-nav")

	if el == nil {
		t.Errorf("GetElementByClassname returned incorrect nil")
	}
}

func TestGetElementByClassnameInvalid(t *testing.T) {
	el := GetElementByClassname(doc, "header-links invalid-class")

	if el != nil {
		t.Errorf("GetElementByClassname returned incorrect value.")
	}
}

func TestGetElementsByClassname(t *testing.T) {
	n := GetElementsByClassname(doc, "section")

	if n == nil {
		t.Errorf("GetElementsByClassname returned incorrect nil")
	}

	if len(n) != 2 {
		t.Errorf("GetElementsByClassname did not return 2 nodes")
	}
}
