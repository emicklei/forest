package rat

import "testing"

func TestExpectStatus404(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/status/404"))
	ExpectStatus(t, r, 404)
}

func TestExpectResponseHeader(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/jsondoc"))
	ExpectHeader(t, r, "Content-Type", "application/json")
}

func TestExpectJsonHash(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/jsondoc"))
	ExpectJsonHash(t, r, func(hash map[string]interface{}) {
		if hash["Value"] != float64(42) {
			t.Errorf("expected 42 but got %v (%T)", hash["Value"], hash["Value"])
		}
	})
}

func TestExpectJsonArray(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/jsonarray"))
	ExpectJsonArray(t, r, func(a []interface{}) {
		if len(a) != 1 && a[0] != 42 {
			t.Errorf("expected 42 but got %v (%T)", a, a)
		}
	})
}
