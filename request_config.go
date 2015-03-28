package rat

import (
	"io"
	"strings"
)

// RequestConfig holds additional information to construct a Http request.
type RequestConfig struct {
	BodyReader io.Reader
	Headers    map[string]string
}

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		Headers: map[string]string{},
	}
}

func (r *RequestConfig) Header(name, value string) *RequestConfig {
	r.Headers[name] = value
	return r
}

func (r *RequestConfig) Body(body string) *RequestConfig {
	r.BodyReader = strings.NewReader(body)
	return r
}
