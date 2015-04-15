package main

import "html/template"

var setup = template.Must(template.New("setup").Parse(`package main

import (
	"net/http"

	. "github.com/emicklei/rat"
)

var api = NewClient("{{.}}", new(http.Client))
`))
