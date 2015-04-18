package rat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
)

var ErrorMessagePrefix = "\n:-( "

// ExpectStatus inspects the response status code.
// If the value is not expected, the complete request, response is logged (iff verbose).
// Return true if the status is as expected.
func ExpectStatus(t T, r *http.Response, status int) bool {
	if r == nil {
		t.Logf("%sExpectStatus: got nil but want Http response", ErrorMessagePrefix)
		return false
	}
	if r.StatusCode != status {
		t.Errorf("%ExpectStatus: got status %d but want %d, %s %v", ErrorMessagePrefix, r.StatusCode, status, r.Request.Method, r.Request.URL)
		if testing.Verbose() {
			Dump(t, r)
		}
		return false
	}
	return true
}

// CheckError simply tests the error and fail is not undefined.
// This is implicity called after sending a Http request.
func CheckError(t T, err error) {
	if err != nil {
		t.Errorf("%sCheckError: did not expect to receive err: %v", ErrorMessagePrefix, err)
	}
}

// ExpectHeader inspects the header of the response.
func ExpectHeader(t T, r *http.Response, name, value string) {
	if r == nil {
		t.Errorf("%sExpectHeader: got nil but want a Http response", ErrorMessagePrefix)
	}
	rname := r.Header.Get(name)
	if rname != value {
		t.Errorf("%sExpectHeader: got header %s=%s but want %s", ErrorMessagePrefix, name, rname, value)
	}
}

// ExpectJSONHash tries to unmarshal the response body into a Go map callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJSONHash(t T, r *http.Response, callback func(hash map[string]interface{})) {
	if r == nil {
		t.Errorf("%sExpectJSONHash: no response available", ErrorMessagePrefix)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("%sExpectJSONHash: unable to read response body:%v", ErrorMessagePrefix, err)
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	dict := map[string]interface{}{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		t.Errorf("%sExpectJSONHash: unable to unmarshal Json:%v", ErrorMessagePrefix, err)
		return
	}
	callback(dict)
}

// ExpectJSONArray tries to unmarshal the response body into a Go slice callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJSONArray(t T, r *http.Response, callback func(array []interface{})) {
	if r == nil {
		t.Errorf("%sExpectJSONArray: no response available", ErrorMessagePrefix)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("%sExpectJSONArray: unable to read response body:%v", ErrorMessagePrefix, err)
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	slice := []interface{}{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		t.Errorf("%ExpectJSONArray: unable to unmarshal Json:%v", ErrorMessagePrefix, err)
	}
	callback(slice)
}

// ExpectString reads the response body into a Go string callback parameter.
// Fail if the body could not be read or unmarshalled.
func ExpectString(t T, r *http.Response, callback func(content string)) {
	if r == nil {
		t.Errorf("%sExpectString: no response available", ErrorMessagePrefix)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("%sExpectString: unable to read response body:%v", ErrorMessagePrefix, err)
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	callback(string(data))
}

// ExpectXMLDocument tries to unmarshal the response body into fields of the provided document (struct).
// Fail if the body could not be read or unmarshalled.
func ExpectXMLDocument(t T, r *http.Response, doc interface{}) {
	if r == nil {
		t.Errorf("%sExpectXMLDocument: no response available", ErrorMessagePrefix)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("%sExpectXMLDocument: unable to read response body:%v", ErrorMessagePrefix, err)
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	err = xml.Unmarshal(data, doc)
	if err != nil {
		t.Errorf("%sExpectXMLDocument: unable to unmarshal Xml:%v", ErrorMessagePrefix, err)
	}
}

// ExpectJSONDocument tries to unmarshal the response body into fields of the provided document (struct).
// Fail if the body could not be read or unmarshalled.
func ExpectJSONDocument(t T, r *http.Response, doc interface{}) {
	if r == nil {
		t.Errorf("%sExpectJSONDocument: no response available", ErrorMessagePrefix)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("%sExpectJSONDocument: unable to read response body :%v", ErrorMessagePrefix, err)
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	err = json.Unmarshal(data, doc)
	if err != nil {
		t.Errorf("%sExpectJSONDocument: unable to unmarshal Json:%v", ErrorMessagePrefix, err)
	}
}
