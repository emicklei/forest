package rat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// RequestConfig holds additional information to construct a Http request.
type RequestConfig struct {
	BodyReader io.Reader
	HeaderMap  http.Header
}

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		HeaderMap: http.Header{},
	}
}

func (r *RequestConfig) Header(name, value string) *RequestConfig {
	r.HeaderMap.Add(name, value)
	return r
}

// Body set the playload as is. No content type is set.
func (r *RequestConfig) Body(body string) *RequestConfig {
	r.BodyReader = strings.NewReader(body)
	return r
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
	}
	r.Body(fmt.Sprintf("cannot encode payload, unknown content type:%s", contentType))
	return r
}
