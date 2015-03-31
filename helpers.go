package rat

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Dump is a convenient method to log the full contents of a request and its response.
func Dump(t T, resp *http.Response) {
	// dump request
	var buffer bytes.Buffer
	buffer.WriteString("\n\n")
	buffer.WriteString(fmt.Sprintf("%v %v\n", resp.Request.Method, resp.Request.URL))
	for k, v := range resp.Request.Header {
		buffer.WriteString(fmt.Sprintf("%s : %v\n", k, strings.Join(v, ",")))
	}

	// dump response
	buffer.WriteString("\n")
	if resp == nil {
		return
	}
	for k, v := range resp.Header {
		buffer.WriteString(fmt.Sprintf("%s : %v\n", k, strings.Join(v, ",")))
	}
	if resp.Body != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			buffer.WriteString(fmt.Sprintf("unable to read body:%v", err))
		} else {
			buffer.WriteString(string(body))
		}
		resp.Body.Close()
	}
	buffer.WriteString("\n-\n")
	t.Logf(buffer.String())
}

func headersString(h http.Header) string {
	if len(h) == 0 {
		return ""
	} else {
		return fmt.Sprintf("%v", h)
	}
}

func copyHeaders(from, to http.Header) {
	for k, list := range from {
		for _, v := range list {
			to.Set(k, v)
		}
	}
}
