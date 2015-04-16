package rat

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"gopkg.in/xmlpath.v2"
)

func XMLPath(t *testing.T, r *http.Response, xpath string) interface{} {
	if r == nil {
		t.Error("no response to read body from")
		return nil
	}
	if r.Body == nil {
		t.Error("no response body to read")
		return nil
	}
	path := xmlpath.MustCompile(xpath)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("unable to read response body")
		return nil
	}
	root, err := xmlpath.Parse(bytes.NewReader(data))
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	if err != nil {
		t.Errorf("unable to parse xml:%v", err)
		return nil
	}
	if value, ok := path.String(root); ok {
		return value
	}
	t.Errorf("no value for path:%s", xpath)
	return nil
}
