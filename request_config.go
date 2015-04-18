package rat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// RequestConfig holds additional information to construct a Http request.
type RequestConfig struct {
	Uri        string
	BodyReader io.Reader
	HeaderMap  http.Header
	Values     url.Values
}

// NewConfig returns a new RequestConfig with initialized empty headers and query parameters.
func NewConfig(pathTemplate string, pathParams ...interface{}) *RequestConfig {
	cfg := &RequestConfig{
		HeaderMap: http.Header{},
		Values:    url.Values{},
	}
	cfg.Path(pathTemplate, pathParams...)
	return cfg
}

// Do calls the one-argument function parameter with the receiver.
// This allows for custom convenience functions without breaking the fluent style.
func (r *RequestConfig) Do(block func(config *RequestConfig)) *RequestConfig {
	block(r)
	return r
}

// Path sets the URL path with optional path parameters.
// format example: /v1/persons/{param}/ + 42 => /v1/persons/42
func (r *RequestConfig) Path(pathTemplate string, pathParams ...interface{}) *RequestConfig {
	var uri bytes.Buffer
	p := 0
	tokens := strings.Split(pathTemplate, "/")
	for _, each := range tokens {
		if len(each) == 0 {
			continue
		}
		uri.WriteString("/")
		if strings.HasPrefix(each, "{") && strings.HasSuffix(each, "}") {
			if p == len(pathParams) {
				// abort
				r.Uri = pathTemplate
				return r
			}
			param := fmt.Sprintf("%v", pathParams[p])
			uri.WriteString(url.QueryEscape(param))
			p++
		} else {
			uri.WriteString(each)
		}
	}
	r.Uri = uri.String()
	return r
}

// Query adds a name=value pair to the list of parameters.
func (r *RequestConfig) Query(name string, value interface{}) *RequestConfig {
	r.Values.Add(name, fmt.Sprintf("%v", value))
	return r
}

// Header adds a name=value pair to the list of headers.
func (r *RequestConfig) Header(name, value string) *RequestConfig {
	r.HeaderMap.Add(name, value)
	return r
}

// Body set the playload as is. No content type is set.
func (r *RequestConfig) Body(body string) *RequestConfig {
	r.BodyReader = strings.NewReader(body)
	return r
}

func (r *RequestConfig) pathAndQuery() string {
	return path.Join(r.Uri, r.Values.Encode())
}

// Content encodes the payload conform the content type given.
func (r *RequestConfig) Content(payload interface{}, contentType string) *RequestConfig {
	r.Header("Content-Type", contentType)
	if strings.Index(contentType, "application/json") != -1 {
		data, err := json.Marshal(payload)
		if err != nil {
			r.Body(fmt.Sprintf("json marshal failed:%v", err))
			return r
		}
		r.BodyReader = bytes.NewReader(data)
		return r
	}
	if strings.Index(contentType, "application/xml") != -1 {
		data, err := xml.Marshal(payload)
		if err != nil {
			r.Body(fmt.Sprintf("xml marshal failed:%v", err))
			return r
		}
		r.BodyReader = bytes.NewReader(data)
		return r
	}
	if strings.Index(contentType, "text/plain") != -1 {
		content, ok := payload.(string)
		if !ok {
			r.Body(fmt.Sprintf("content is not a string:%v", payload))
			return r
		}
		r.BodyReader = strings.NewReader(content)
		return r
	}
	bits, ok := payload.([]byte)
	if ok {
		r.BodyReader = bytes.NewReader(bits)
		return r
	}
	r.Body(fmt.Sprintf("cannot encode payload, unknown content type:%s", contentType))
	return r
}
