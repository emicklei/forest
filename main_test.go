package rat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// go test -coverprofile=cover.out && go tool cover -html=cover.out

var tsAPI *APITesting

func TestMain(m *testing.M) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "404") {
			w.WriteHeader(404)
			return
		}
		if strings.HasSuffix(r.URL.Path, "jsonarray") {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintln(w, "[42]")
			return
		}
		if strings.HasSuffix(r.URL.Path, "jsondoc") {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintln(w, "{\"Value\":42}")
			return
		}
		if strings.HasSuffix(r.URL.Path, "json-nested-doc") {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintln(w, `{"Root": {"Child":12} }`)
			return
		}
		if strings.HasSuffix(r.URL.Path, "xmldoc") {
			w.Header().Add("Content-Type", "application/xml")
			fmt.Fprintln(w, `<?xml version="1.0"?><Root><Child name="answer"><Value>42</Value></Child></Root>`)
			return
		}
		if strings.HasSuffix(r.URL.Path, "echo") {
			w.Header().Add("Content-Type", "application/octet-stream")
			w.Header().Add("ECHO", r.Header.Get("ECHO"))
			data, _ := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			w.Write(data)
			return
		}
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" {
			w.WriteHeader(204)
		}
		// 200 is written
	}))
	tsAPI = NewClient(ts.URL, new(http.Client))

	exitCode := m.Run()
	// on early exit close will not be called
	ts.Close()
	os.Exit(exitCode)
}
