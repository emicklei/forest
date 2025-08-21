package forest

import (
	"net/http"
)

// APITesting provides functions to call a REST api and validate its responses.
type APITesting struct {
	BaseURL string
	client  *http.Client
}

// NewClient returns a new ApiTesting that can be used to send Http requests.
func NewClient(baseURL string, httpClient *http.Client) *APITesting {
	return &APITesting{
		BaseURL: baseURL,
		client:  httpClient,
	}
}

// GET sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
// GET sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *APITesting) GET(t T, config *RequestConfig) *http.Response {
	t.Helper()
	return a.doRequest(t, http.MethodGet, config)
}

// POST sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *APITesting) POST(t T, config *RequestConfig) *http.Response {
	t.Helper()
	return a.doRequest(t, http.MethodPost, config)
}

// PUT sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *APITesting) PUT(t T, config *RequestConfig) *http.Response {
	t.Helper()
	return a.doRequest(t, http.MethodPut, config)
}

// DELETE sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *APITesting) DELETE(t T, config *RequestConfig) *http.Response {
	t.Helper()
	return a.doRequest(t, http.MethodDelete, config)
}

// PATCH sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *APITesting) PATCH(t T, config *RequestConfig) *http.Response {
	t.Helper()
	return a.doRequest(t, http.MethodPatch, config)
}

func (a *APITesting) doRequest(t T, method string, config *RequestConfig) *http.Response {
	t.Helper()
	httpReq, err := config.Build(method, a.BaseURL)
	if err != nil {
		logfatal(t, sfatalf("%s: invalid Url:%s", method, a.BaseURL+config.pathAndQuery()))
	}
	if config.logRequestLine {
		Logf(t, "\n%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
	}
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return ensureResponse(httpReq, resp)
}

// Do sends a Http request using a Http method (GET,PUT,POST,....) and config (headers,...)
// The request is not logged and any URL build error or send error will be returned.
func (a *APITesting) Do(method string, config *RequestConfig) (*http.Response, error) {
	httpReq, err := config.Build(method, a.BaseURL)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(httpReq)
	return ensureResponse(httpReq, resp), err
}
