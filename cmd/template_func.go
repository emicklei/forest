package main

import "html/template"

var fun = template.Must(template.New("fun").Parse(`package main

import (
	"net/http"
	"testing"

	. "github.com/emicklei/rat"
)

func Test_{{.Name}}(t *testing.T) {
	cfg := NewConfig("{{.Path}}")
	r := api.{{.Method}}(t, cfg)
	ExpectStatus(t, r, {{.Status}})
}
`))

var fun_pathparams = template.Must(template.New("fun_pathparams").Parse(`package main

import (
	"net/http"
	"testing"

	. "github.com/emicklei/rat"
)

func Test_{{.Name}}(t *testing.T) {
	cfg := NewConfig("").Path("{{.Path}}","?")
	r := api.{{.Method}}(t, cfg)
	ExpectStatus(t, r, {{.Status}})
}
`))
