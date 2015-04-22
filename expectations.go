package rat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
)

// ExpectStatus inspects the response status code.
// If the value is not expected, the complete request, response is logged (iff verbose).
// Return true if the status is as expected.
func ExpectStatus(t T, r *http.Response, status int) bool {
	if r == nil {
		t.Error(serrorf("ExpectStatus: got nil but want Http response"))
		return false
	}
	if r.StatusCode != status {
		t.Error(serrorf("ExpectStatus: got status %d but want %d, %s %v", r.StatusCode, status, r.Request.Method, r.Request.URL))
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
		t.Error(serrorf("CheckError: did not expect to receive err: %v", err))
	}
}

// ExpectHeader inspects the header of the response.
func ExpectHeader(t T, r *http.Response, name, value string) {
	if r == nil {
		t.Error(serrorf("ExpectHeader: got nil but want a Http response"))
	}
	rname := r.Header.Get(name)
	if rname != value {
		t.Error(serrorf("ExpectHeader: got header %s=%s but want %s", name, rname, value))
	}
}

// ExpectJSONHash tries to unmarshal the response body into a Go map callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJSONHash(t T, r *http.Response, callback func(hash map[string]interface{})) {
	if r == nil {
		t.Error(serrorf("ExpectJSONHash: no response available"))
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Error(serrorf("ExpectJSONHash: unable to read response body:%v", err))
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	dict := map[string]interface{}{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		t.Error(serrorf("ExpectJSONHash: unable to unmarshal Json:%v", err))
		return
	}
	callback(dict)
}

// ExpectJSONArray tries to unmarshal the response body into a Go slice callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJSONArray(t T, r *http.Response, callback func(array []interface{})) {
	if r == nil {
		t.Error(serrorf("ExpectJSONArray: no response available"))
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Error(serrorf("ExpectJSONArray: unable to read response body:%v", err))
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	slice := []interface{}{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		t.Error(serrorf("ExpectJSONArray: unable to unmarshal Json:%v", err))
	}
	callback(slice)
}

// ExpectString reads the response body into a Go string callback parameter.
// Fail if the body could not be read or unmarshalled.
func ExpectString(t T, r *http.Response, callback func(content string)) {
	if r == nil {
		t.Error(serrorf("ExpectString: no response available"))
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Error(serrorf("ExpectString: unable to read response body:%v", err))
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
		t.Error(serrorf("ExpectXMLDocument: no response available"))
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Error(serrorf("ExpectXMLDocument: unable to read response body:%v", err))
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	err = xml.Unmarshal(data, doc)
	if err != nil {
		t.Error(serrorf("ExpectXMLDocument: unable to unmarshal Xml:%v", err))
	}
}

// ExpectJSONDocument tries to unmarshal the response body into fields of the provided document (struct).
// Fail if the body could not be read or unmarshalled.
func ExpectJSONDocument(t T, r *http.Response, doc interface{}) {
	if r == nil {
		t.Error(serrorf("ExpectJSONDocument: no response available"))
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Error(serrorf("ExpectJSONDocument: unable to read response body :%v", err))
		return
	}
	// put the body back for re-reads
	r.Body = &closeableReader{bytes.NewReader(data)}

	err = json.Unmarshal(data, doc)
	if err != nil {
		t.Error(serrorf("ExpectJSONDocument: unable to unmarshal Json:%v", err))
	}
}
