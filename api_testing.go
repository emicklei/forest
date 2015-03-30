package rat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ApiTesting provides functions to call a REST api and validate its responses.
type ApiTesting struct {
	BaseUrl string
	client  *http.Client
}

// NewClient returns a new ApiTesting that can be used to send Http requests.
func NewClient(baseUrl string, httpClient *http.Client) *ApiTesting {
	return &ApiTesting{
		BaseUrl: baseUrl,
		client:  httpClient,
	}
}

// NewConfig returns a new RequestConfig.
func (a *ApiTesting) NewConfig(staticPath string) *RequestConfig {
	return NewConfig(staticPath)
}

// GET sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) GET(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("GET", a.BaseUrl+config.Uri, nil)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, httpReq.Header)
	resp, err := a.client.Do(httpReq)
	a.CheckError(t, err)
	return resp
}

// POST sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) POST(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("POST", a.BaseUrl+config.Uri, config.BodyReader)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, httpReq.Header)
	resp, err := a.client.Do(httpReq)
	a.CheckError(t, err)
	return resp
}

// PUT sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) PUT(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("PUT", a.BaseUrl+config.Uri, config.BodyReader)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, httpReq.Header)
	resp, err := a.client.Do(httpReq)
	a.CheckError(t, err)
	return resp
}

// DELETE sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) DELETE(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("DELETE", a.BaseUrl+config.Uri, nil)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	t.Logf("%v %v %v", httpReq.Method, httpReq.URL, httpReq.Header)
	resp, err := a.client.Do(httpReq)
	a.CheckError(t, err)
	return resp
}

// CheckError simply tests the error and fail is not undefined.
// This is implicity called after sending a Http request.
func (a *ApiTesting) CheckError(t T, err error) {
	if err != nil {
		t.Fatalf("did not expect to receive err: %v", err)
	}
}

// ExpectStatus inspects the response status code.
// If the value is not expected, the body (if any) is logged.
func (a ApiTesting) ExpectStatus(t T, r *http.Response, status int) {
	if r == nil {
		t.Fatalf("expected response but got nil")
	}
	if r.StatusCode != status {
		t.Errorf("expected status %d but got %d", status, r.StatusCode)
		data, _ := ioutil.ReadAll(r.Body)
		t.Logf("%s", string(data))
	}
}

// ExpectHeader inspects the header of the response.
func (a ApiTesting) ExpectHeader(t T, r *http.Response, name, value string) {
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
func (a ApiTesting) ExpectJsonHash(t T, r *http.Response, callback func(hash map[string]interface{})) {
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
func (a ApiTesting) ExpectJsonArray(t T, r *http.Response, callback func(array []interface{})) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("unable to read response body:%v", err)
	}
	defer r.Body.Close()

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
func (a ApiTesting) ExpectString(t T, r *http.Response, callback func(content string)) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("unable to read response body:%v", err)
	}
	defer r.Body.Close()

	callback(string(data))
}

// Dump is a convenient method to log the full contents of a response
func (a ApiTesting) Dump(t T, resp *http.Response) {
	if resp == nil {
		t.Errorf("no response")
		return
	}
	if resp.ContentLength == 0 {
		t.Logf("empty response")
		return
	}
	for k, v := range resp.Header {
		t.Logf("%s : %v", k, v)
	}
	if resp.Body != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf(string(body))
		resp.Body.Close()
	}
}

func copyHeaders(from, to http.Header) {
	for k, list := range from {
		for _, v := range list {
			to.Set(k, v)
		}
	}
}
