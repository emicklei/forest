package forest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// maskHeaderNames is used to prevent logging secrets.
var maskHeaderNames = []string{}

// MaskChar is used the create a masked header value.
var MaskChar = "*"

// MaskingFunc is the function used to mask a header value.
// If nil then the default masking is applied.
var MaskingFunc func(value string) string

// MaskHeader is used to prevent logging secrets.
func MaskHeader(name string) {
	maskHeaderNames = append(maskHeaderNames, name)
}

// IsMaskedHeader return true if the name is part of (case-insensitive match) the MaskHeaderNames.
func IsMaskedHeader(name string) bool {
	for _, each := range maskHeaderNames {
		if strings.ToLower(each) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// CookieNamed returns the cookie with matching name. Returns nil if not found.
func CookieNamed(resp *http.Response, name string) *http.Cookie {
	for _, each := range resp.Cookies() {
		if each.Name == name {
			return each
		}
	}
	return nil
}

// DumpTemplate is the template used to dump request and response information.
// If nil then the default dump format is used.
var DumpTemplate *template.Template

// DumpData is the structure passed to the DumpTemplate.
type DumpData struct {
	Request  *http.Request
	Response *http.Response
	Body     string
}

// Dump is a convenient method to log the full contents of a request and its response.
func Dump(t T, resp *http.Response) {
	if DumpTemplate != nil {
		body, err := ReadAndRestoreBody(resp)
		if err != nil {
			// ignore error, dump what we have
		}
		var buffer bytes.Buffer
		err = DumpTemplate.Execute(&buffer, DumpData{
			Request:  resp.Request,
			Response: resp,
			Body:     string(body),
		})
		if err != nil {
			logerror(t, serrorf("dump template failed:%v", err))
			return
		}
		Logf(t, buffer.String())
		return
	}
	// default dump
	// dump request
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%v %v\n", resp.Request.Method, resp.Request.URL))
	for k, v := range resp.Request.Header {
		if IsMaskedHeader(k) && len(v) > 0 {
			v = []string{maskedHeaderValue(v[0])}
		}
		if len(k) > 0 {
			buffer.WriteString(fmt.Sprintf("%s : %v\n", k, strings.Join(v, ",")))
		}
	}
	// dump request payload, only available is there is a Body.
	if resp != nil && resp.Request != nil && resp.Request.Body != nil {
		rc, _ := resp.Request.GetBody()
		body, err := io.ReadAll(rc)
		if err != nil {
			buffer.WriteString(fmt.Sprintf("unable to read request body:%v", err))
		} else {
			if len(body) > 0 {
				buffer.WriteString("\n")
			}
			buffer.WriteString(string(body))
			buffer.WriteString("\n")
		}
	}
	// dump response payload
	if resp == nil {
		buffer.WriteString("-- no response --")
		Logf(t, buffer.String())
		return
	}
	buffer.WriteString(fmt.Sprintf("\n%s\n", resp.Status))
	for k, v := range resp.Header {
		if len(k) > 0 {
			buffer.WriteString(fmt.Sprintf("%s : %v\n", k, strings.Join(v, ",")))
		}
	}
	if resp.Body != nil {
		body, err := ReadAndRestoreBody(resp)
		if err != nil {
			if resp.StatusCode/100 == 3 {
				// redirect closes body ; nothing to read
				buffer.WriteString("\n")
			} else {
				buffer.WriteString(fmt.Sprintf("unable to read response body:%v", err))
			}
		} else {
			if len(body) > 0 {
				buffer.WriteString("\n")
			}
			buffer.WriteString(string(body))
		}
	}
	buffer.WriteString("\n")
	Logf(t, "%s", buffer.String())
}

// ReadAndRestoreBody reads the content of the response body and puts it back for re-reads.
// Be aware that this is not a memory-friendly operation.
func ReadAndRestoreBody(r *http.Response) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, err
	}
	// put the body back for re-reads
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	return data, nil
}

type skippeable interface {
	Skipf(string, ...interface{})
}

// SkipUnless will Skip the test unless the LABELS environment variable includes any of the provided labels.
//
//	LABELS=integration,nightly go test -v
func SkipUnless(t skippeable, labels ...string) {
	env := strings.Split(os.Getenv("LABELS"), ",")
	for _, each := range labels {
		for _, other := range env {
			if each == other {
				return
			}
		}
	}
	t.Skipf("skipped because provided LABELS=%v does not include any of %v", env, labels)
}

func headersString(h http.Header) string {
	if len(h) == 0 {
		return ""
	}
	masked := http.Header{}
	for k, v := range h {
		if IsMaskedHeader(k) && len(v) > 0 {
			v = []string{maskedHeaderValue(v[0])}
		}
		masked[k] = v
	}
	return fmt.Sprintf("%v", masked)
}

func maskedHeaderValue(s string) string {
	if MaskingFunc != nil {
		return MaskingFunc(s)
	}
	stars := strings.Repeat(MaskChar, 3)
	return fmt.Sprintf("%s(masked %d chars)%s", stars, len(s), stars)
}

func copyHeaders(from, to http.Header) {
	for k, list := range from {
		for _, v := range list {
			to.Set(k, v)
		}
	}
}

func setFormData(config *RequestConfig, req *http.Request) {
	// set form data if available
	if len(config.FormData) > 0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Body = ioutil.NopCloser(strings.NewReader(config.FormData.Encode()))
	}
}

func setBasicAuth(config *RequestConfig, req *http.Request) {
	if len(config.User) > 0 {
		req.SetBasicAuth(config.User, config.Password)
	}
}

func ensureResponse(req *http.Request, resp *http.Response) *http.Response {
	if resp != nil {
		return resp
	}
	// wrap the request into an empty response ; Body must be nil
	wrap := new(http.Response)
	wrap.Request = req
	return wrap
}

func URLPathEncode(path string) string {
	buf := new(bytes.Buffer)
	for _, each := range path {
		switch each {
		case '+':
			buf.WriteString("%20")
		case ' ':
			buf.WriteString("%20")
		default:
			buf.WriteRune(each)
		}
	}
	return buf.String()
}

func setCookies(config *RequestConfig, req *http.Request) {
	for _, each := range config.Cookies {
		req.AddCookie(each)
	}
}
