package forest

import (
	"testing"
)

func TestJSONPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/json-nested-doc"))
	v, ok := JSONPath(t, r, "$.Root.Child")
	if !ok || v != 12.0 {
		t.Errorf("got %v (%T) want 12.0", v, v)
	}
}

func TestJSONArrayPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/json-array-of-doc"))
	v, ok := JSONArrayPath(t, r, "$[1].digit")
	if !ok || v != 2.0 {
		t.Errorf("got %v (%T) want 2.0", v, v)
	}
}

func TestMustJSONPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/json-nested-doc"))
	if v := MustJSONPath(t, r, "$.Root.Child"); v != 12.0 {
		t.Errorf("got %v (%T) want 12.0", v, v)
	}
}

func TestMustJSONArrayPath(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/json-array-of-doc"))
	if v := MustJSONArrayPath(t, r, "$[1].digit"); v != 2.0 {
		t.Errorf("got %v (%T) want 2.0", v, v)
	}
}

func TestMustJSONPathPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	r := tsAPI.GET(t, NewConfig("/json-nested-doc"))
	MustJSONPath(t, r, "$.Root.Invalid")
}

func TestMustJSONArrayPathPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	r := tsAPI.GET(t, NewConfig("/json-array-of-doc"))
	MustJSONArrayPath(t, r, "$[2].digit")
}
