package rat

import (
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
		t.Logf("ExpectStatus failed: got nil but want Http response")
		return false
	}
	if r.StatusCode != status {
		t.Errorf("ExpectStatus failed: got status %d but want %d, %s %v", r.StatusCode, status, r.Request.Method, r.Request.URL)
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
		t.Errorf("CheckError failed: did not expect to receive err: %v", err)
	}
}

// ExpectHeader inspects the header of the response.
func ExpectHeader(t T, r *http.Response, name, value string) {
	if r == nil {
		t.Errorf("ExpectHeader failed: got nil but want a Http response")
	}
	rname := r.Header.Get(name)
	if rname != value {
		t.Errorf("ExpectHeader failed: got header %s=%s but want %s", name, rname, value)
	}
}

// ExpectJsonHash tries to unmarshal the response body into a Go map callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJsonHash(t T, r *http.Response, callback func(hash map[string]interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("ExpectJsonHash failed: unable to read response body:%v", err)
		return
	}

	dict := map[string]interface{}{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		t.Errorf("ExpectJsonHash failed: unable to unmarshal Json:%v", err)
	}
	callback(dict)
}

// ExpectJsonArray tries to unmarshal the response body into a Go slice callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJsonArray(t T, r *http.Response, callback func(array []interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("ExpectJsonArray failed: unable to read response body:%v", err)
		return
	}

	t.Logf("%s", string(data))
	slice := []interface{}{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		t.Errorf("ExpectJsonArray failed: unable to unmarshal Json:%v", err)
	}
	callback(slice)
}

// ExpectString reads the response body into a Go string callback parameter.
// Fail if the body could not be read or unmarshalled.
func ExpectString(t T, r *http.Response, callback func(content string)) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("ExpectString failed: unable to read response body:%v", err)
		return
	}

	callback(string(data))
}

// ExpectXmlDocument tries to unmarshal the response body into field of the provided document.
// Fail if the body could not be read or unmarshalled.
func ExpectXmlDocument(t T, r *http.Response, doc interface{}) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Errorf("ExpectXmlDocument failed: unable to read response body:%v", err)
		return
	}
	t.Logf("%s", string(data))

	err = xml.Unmarshal(data, doc)
	if err != nil {
		t.Errorf("ExpectXmlDocument failed: unable to unmarshal Xml:%v", err)
	}
}
