package forest

import (
	"net/http"
	"testing"
)

func TestXMLPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/xmldoc"))
	v := XMLPath(t, r, "/Root/Child/Value")
	if v != "1" {
		t.Errorf("got %v but want 1", v)
	}
}

// TODO
func TestXMLPathAllChilds(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/xmldoc"))
	m := new(mockedT)
	kids := XMLPath(m, r, "Root//Child")
	t.Logf("%#v", kids)
}

func TestXMLPathWrongPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/xmldoc"))
	m := new(mockedT)
	XMLPath(m, r, "/Root/Child/ValueX")
	if m.logMessage != "XMLPath: no value for path: /Root/Child/ValueX" {
		t.Errorf("expected other message than:[%s]", m.errorMessage)
	}
}

func TestXMLPathInvalidPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/xmldoc"))
	m := new(mockedT)
	XMLPath(m, r, "/{Root}")
	if m.logMessage != "XMLPath: invalid xpath expression:compiling xml path \"/{Root}\":1: missing name" {
		t.Errorf("expected other message than:[%s]", m.errorMessage)
	}
}

func TestXMLPathNoResponse(t *testing.T) {
	m := new(mockedT)
	XMLPath(m, nil, "/Root")
	if m.logMessage != "XMLPath: no response to read body from" {
		t.Errorf("expected other message than:[%s]", m.fatalMessage)
	}
}

func TestXMLPathNoBody(t *testing.T) {
	m := new(mockedT)
	r := new(http.Response)
	XMLPath(m, r, "/Root")
	if m.logMessage != "XMLPath: no response body to read" {
		t.Errorf("expected other message than:[%s]", m.fatalMessage)
	}
}

func TestXMLPathNoDocument(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/404"))
	m := new(mockedT)
	XMLPath(m, r, "/Root")
	if m.logMessage != "XMLPath: no value for path: /Root" {
		t.Errorf("expected other message than:[%s]", m.errorMessage)
	}
}
