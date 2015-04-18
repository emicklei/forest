package rat

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

// XMLPath returns the value found by following the xpath expression in a XML document (payload of response).
func XMLPath(t T, r *http.Response, xpath string) interface{} {
	if r == nil {
		t.Fatalf("%sXMLPath: no response to read body from", ErrorMessagePrefix)
		return nil
	}
	if r.Body == nil {
		t.Fatalf("%sXMLPath: no response body to read", ErrorMessagePrefix)
		return nil
	}
	path, err := xmlpath.Compile(xpath)
	if err != nil {
		t.Errorf("%sXMLPath: invalid xpath expression:%v", ErrorMessagePrefix, err)
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("%sXMLPath: unable to read response body", ErrorMessagePrefix)
		return nil
	}
	root, err := xmlpath.Parse(bytes.NewReader(data))
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	if err != nil {
		t.Errorf("%sXMLPath: unable to parse xml:%v", ErrorMessagePrefix, err)
		return nil
	}
	if value, ok := path.String(root); ok {
		return value
	}
	t.Errorf("%sXMLPath: no value for path: %s", ErrorMessagePrefix, xpath)
	return nil
}
