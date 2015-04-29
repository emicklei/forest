package forest

import "testing"

func TestExpectStatus404(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/status/404"))
	ExpectStatus(t, r, 404)
}

func TestExpectResponseHeader(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/jsondoc"))
	ExpectHeader(t, r, "Content-Type", "application/json")
}

func TestExpectJSONHash(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/jsondoc"))
	ExpectJSONHash(t, r, func(hash map[string]interface{}) {
		if hash["Value"] != float64(42) {
			t.Error(serrorf("expected 42 but got %v (%T)", hash["Value"], hash["Value"]))
		}
	})
}

func TestExpectJSONArray(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/jsonarray"))
	ExpectJSONArray(t, r, func(a []interface{}) {
		if len(a) != 1 && a[0] != 42 {
			t.Errorf("expected 42 but got %v (%T)", a, a)
		}
	})
}

type Root struct {
	Child struct {
		Name  string `xml:"name,attr"`
		Value int
	}
}

func TestExpectXMLDoc(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/xmldoc"))
	var root Root
	ExpectXMLDocument(t, r, &root)
	if root.Child.Name != "answer" {
		t.Error("expected attribute was the answer")
	}
}

func TestExpectJSONDoc(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/jsondoc"))
	root := struct {
		Value int
	}{}
	ExpectJSONDocument(t, r, &root)
	if root.Value != 42 {
		t.Errorf("expected 42 was the answer, got %v", root.Value)
	}
}
