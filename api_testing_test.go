package rat

import "testing"

func TestExpectStatus404(t *testing.T) {
	r := tsApi.GET(t, "/status/404", NewRequestConfig())
	tsApi.ExpectStatus(t, r, 404)
}

func TestExpectResponseHeader(t *testing.T) {
	r := tsApi.GET(t, "/jsondoc", NewRequestConfig())
	tsApi.ExpectHeader(t, r, "Content-Type", "application/json")
}

func TestExpectJsonHash(t *testing.T) {
	r := tsApi.GET(t, "/jsondoc", NewRequestConfig())
	tsApi.ExpectJsonHash(t, r, func(hash map[string]interface{}) {
		if hash["Value"] != float64(42) {
			t.Errorf("expected 42 but got %v (%T)", hash["Value"], hash["Value"])
		}
	})
}

func TestExpectJsonArray(t *testing.T) {
	r := tsApi.GET(t, "/jsonarray", NewRequestConfig())
	tsApi.ExpectJsonArray(t, r, func(a []interface{}) {
		if len(a) != 1 && a[0] != 42 {
			t.Errorf("expected 42 but got %v (%T)", a, a)
		}
	})
}

func TestDump(t *testing.T) {
	r := tsApi.GET(t, "/jsonarray", NewRequestConfig())
	tsApi.Dump(t, r)
}

func TestPost(t *testing.T) {
	r := tsApi.POST(t, "/echo", NewRequestConfig().Body("data").Header("ECHO", "ping"))
	tsApi.ExpectString(t, r, func(m string) {
		if m != "data" {
			t.Errorf("expected data but got %v", m)
		}
	})
	tsApi.ExpectHeader(t, r, "ECHO", "ping")
}
func TestDelete(t *testing.T) {
	r := tsApi.DELETE(t, "/", NewRequestConfig())
	tsApi.ExpectStatus(t, r, 204)
}

func TestPut(t *testing.T) {
	r := tsApi.PUT(t, "/", NewRequestConfig())
	tsApi.ExpectStatus(t, r, 204)
}
