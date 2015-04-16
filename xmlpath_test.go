package rat

import "testing"

func TestXMLPath(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/xmldoc"))
	v := XMLPath(t, r, "/Root/Child/Value")
	if v != "42" {
		t.Errorf("got %v but want 42", v)
	}
}
