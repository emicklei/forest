package main

import "html/template"

var setup = template.Must(template.New("setup").Parse(`package main

import (
	"net/http"

	. "github.com/emicklei/forest"
)

var api = NewClient("{{.}}", new(http.Client))
`))
