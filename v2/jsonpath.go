package v2

import (
	"encoding/json"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
	"github.com/emicklei/forest"
)

// JSONPath returns the value found by following the dotted path in a JSON document hash.
// E.g. "$.chapters[0].title" in `{ "chapters" : [{"title":"Go a long way"}] }`
func JSONPath(t forest.T, r *http.Response, path string) (interface{}, bool) {
	t.Helper()
	var doc interface{}
	data, err := forest.ReadAndRestoreBody(r)
	if err != nil {
		forest.Logf(t, "unable to read response body:%v", err)
		return nil, false
	}
	err = json.Unmarshal(data, &doc)
	if err != nil {
		forest.Logf(t, "unable to unmarshal json:%v", err)
		return nil, false
	}
	val, err := jsonpath.Get(path, doc)
	if err != nil {
		return nil, false
	}
	return val, true
}

// JSONArrayPath returns the value found by following the dotted path in a JSON array.
// This function is deprecated, use JSONPath instead.
func JSONArrayPath(t forest.T, r *http.Response, dottedPath string) (interface{}, bool) {
	return JSONPath(t, r, dottedPath)
}

// MustJSONPath returns the value found by following the dotted path in a JSON document hash.
// It panics if the path is not found.
func MustJSONPath(t forest.T, r *http.Response, path string) interface{} {
	t.Helper()
	val, ok := JSONPath(t, r, path)
	if !ok {
		panic("JSONPath not found: " + path)
	}
	return val
}

// MustJSONArrayPath returns the value found by following the dotted path in a JSON array.
// This function is deprecated, use MustJSONPath instead.
func MustJSONArrayPath(t forest.T, r *http.Response, dottedPath string) interface{} {
	return MustJSONPath(t, r, dottedPath)
}
