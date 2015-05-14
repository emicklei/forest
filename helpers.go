package forest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Dump is a convenient method to log the full contents of a request and its response.
func Dump(t T, resp *http.Response) {
	// dump request
	var buffer bytes.Buffer
	buffer.WriteString("\n")
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
		// put the body back for re-reads
		resp.Body = &closeableReader{bytes.NewReader(body)}
	}
	buffer.WriteString("\n")
	t.Logf(buffer.String())
}

type skippeable interface {
	Skipf(string, ...interface{})
}

// SkipUnless will Skip the test unless the LABELS environment variable includes any of the provided labels.
//
//	LABELS=integration,nightly go test -v
func SkipUnless(s skippeable, labels ...string) {
	env := strings.Split(os.Getenv("LABELS"), ",")
	for _, each := range labels {
		for _, other := range env {
			if each == other {
				return
			}
		}
	}
	s.Skipf("skipped because provided LABELS=%v does not include any of %v", env, labels)
}

func headersString(h http.Header) string {
	if len(h) == 0 {
		return ""
	}
	return fmt.Sprintf("%v", h)
}

func copyHeaders(from, to http.Header) {
	for k, list := range from {
		for _, v := range list {
			to.Set(k, v)
		}
	}
}
