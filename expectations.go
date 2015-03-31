package rat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

// ExpectStatus inspects the response status code.
// If the value is not expected, the body (if any) is logged.
func ExpectStatus(t T, r *http.Response, status int) {
	if r == nil {
		Logf("got nil but want Http response")
		t.FailNow()
	}
	if r.StatusCode != status {
		Errorf("got status %d but want %d", r.StatusCode, status)
		t.Fail()
		data, _ := ioutil.ReadAll(r.Body)
		if testing.Verbose() {
			Logf("http response body:%s", string(data))
		}
	}
}

// CheckError simply tests the error and fail is not undefined.
// This is implicity called after sending a Http request.
func CheckError(t T, err error) {
	if err != nil {
		Logf("did not expect to receive err: %v", err)
		t.FailNow()
	}
}

// ExpectHeader inspects the header of the response.
func ExpectHeader(t T, r *http.Response, name, value string) {
	if r == nil {
		Logf("got nil but want a Http response")
		t.FailNow()
	}
	rname := r.Header.Get(name)
	if rname != value {
		Logf("got header %s=%s but want %s", name, rname, value)
		t.Fail()
	}
}

// ExpectJsonHash tries to unmarshal the response body into a Go map callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJsonHash(t T, r *http.Response, callback func(hash map[string]interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logf("unable to read response body:%v", err)
		t.FailNow()
	}
	defer r.Body.Close()

	dict := map[string]interface{}{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		if len(data) > 0 {
			Logf("%s", string(data))
		}
		Logf("unable to unmarshal Json:%v", err)
		t.FailNow()
	}
	callback(dict)
}

// ExpectJsonArray tries to unmarshal the response body into a Go slice callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJsonArray(t T, r *http.Response, callback func(array []interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logf("unable to read response body:%v", err)
		t.FailNow()
	}
	defer r.Body.Close()

	if testing.Verbose() {
		Logf("%s", string(data))
	}
	slice := []interface{}{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		Logf("%s", string(data))
		Logf("unable to unmarshal Json:%v", err)
		t.FailNow()
	}
	callback(slice)
}

// ExpectString tries to convert the response body into a Go string callback parameter.
// Fail if the body could not be read.
func ExpectString(t T, r *http.Response, callback func(content string)) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logf("unable to read response body:%v", err)
		t.FailNow()
	}
	defer r.Body.Close()

	callback(string(data))
}
