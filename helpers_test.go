package rat

import "testing"

func TestDump(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/jsonarray"))
	Dump(t, r)
	// check if we can read it again
	Dump(t, r)
}

func TestJSONPath(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/json-nested-doc"))
	if v := JSONPath(t, r, ".Root.Child"); v != 12.0 {
		t.Errorf("got %v (%T) want 12", v, v)
	}
}
