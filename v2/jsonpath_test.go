package v2

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/emicklei/forest"
)

var tsAPI *forest.APITesting

func TestMain(m *testing.M) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "json-nested-doc") {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintln(w, `{"Root": {"Child":12} }`)
			return
		}
		if strings.HasSuffix(r.URL.Path, "json-array-of-doc") {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintln(w, `[{"digit":1}, {"digit":2} ]`)
			return
		}
	}))
	tsAPI = forest.NewClient(ts.URL, new(http.Client))

	exitCode := m.Run()
	ts.Close()
	os.Exit(exitCode)
}

func TestJSONPath(t *testing.T) {
	r := tsAPI.GET(t, forest.NewConfig("/json-nested-doc"))
	v, ok := JSONPath(t, r, "$.Root.Child")
	if !ok || v != 12.0 {
		t.Errorf("got %v (%T) want 12.0", v, v)
	}
}

func TestJSONArrayPath(t *testing.T) {
	r := tsAPI.GET(t, forest.NewConfig("/json-array-of-doc"))
	v, ok := JSONArrayPath(t, r, "$[1].digit")
	if !ok || v != 2.0 {
		t.Errorf("got %v (%T) want 2.0", v, v)
	}
}

func TestMustJSONPath(t *testing.T) {
	r := tsAPI.GET(t, forest.NewConfig("/json-nested-doc"))
	if v := MustJSONPath(t, r, "$.Root.Child"); v != 12.0 {
		t.Errorf("got %v (%T) want 12.0", v, v)
	}
}

func TestMustJSONArrayPath(t *testing.T) {
	r := tsAPI.GET(t, forest.NewConfig("/json-array-of-doc"))
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
	r := tsAPI.GET(t, forest.NewConfig("/json-nested-doc"))
	MustJSONPath(t, r, "$.Root.Invalid")
}

func TestMustJSONArrayPathPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	r := tsAPI.GET(t, forest.NewConfig("/json-array-of-doc"))
	MustJSONArrayPath(t, r, "$[2].digit")
}
