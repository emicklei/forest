package rat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ExpectStatus inspects the response status code.
// If the value is not expected, the body (if any) is logged.
func ExpectStatus(t T, r *http.Response, status int) {
	if r == nil {
		t.Fatalf("expected response but got nil")
	}
	if r.StatusCode != status {
		t.Errorf("expected status %d but got %d", status, r.StatusCode)
		data, _ := ioutil.ReadAll(r.Body)
		t.Logf("response body:%s", string(data))
	}
}

// CheckError simply tests the error and fail is not undefined.
// This is implicity called after sending a Http request.
func CheckError(t T, err error) {
	if err != nil {
		t.Fatalf("did not expect to receive err: %v", err)
	}
}

// ExpectHeader inspects the header of the response.
func ExpectHeader(t T, r *http.Response, name, value string) {
	if r == nil {
		t.Fatalf("expected response but got nil")
	}
	rname := r.Header.Get(name)
	if rname != value {
		t.Errorf("expected header %s=%s but got %s", name, value, rname)
	}
}

// ExpectJsonHash tries to unmarshal the response body into a Go map callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJsonHash(t T, r *http.Response, callback func(hash map[string]interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("unable to read response body:%v", err)
	}
	defer r.Body.Close()

	dict := map[string]interface{}{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		t.Logf("%s", string(data))
		t.Fatalf("unable to unmarshal Json:%v", err)
	}
	callback(dict)
}

// ExpectJsonArray tries to unmarshal the response body into a Go slice callback parameter.
// Fail if the body could not be read or if unmarshalling was not possible.
func ExpectJsonArray(t T, r *http.Response, callback func(array []interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("unable to read response body:%v", err)
	}
	defer r.Body.Close()

	t.Logf(string(data))
	slice := []interface{}{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		t.Logf("%s", string(data))
		t.Fatalf("unable to unmarshal Json:%v", err)
	}
	callback(slice)
}

// ExpectString tries to convert the response body into a Go string callback parameter.
// Fail if the body could not be read.
func ExpectString(t T, r *http.Response, callback func(content string)) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("unable to read response body:%v", err)
	}
	defer r.Body.Close()

	callback(string(data))
}
