# forest - for testing REST api-s in Go

[![Build Status](https://travis-ci.org/emicklei/forest.png)](https://travis-ci.org/emicklei/forest)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/forest)](https://goreportcard.com/report/github.com/emicklei/forest)
[![GoDoc](https://godoc.org/github.com/emicklei/forest?status.svg)](https://godoc.org/github.com/emicklei/forest)

This package provides a few simple helper types and functions to create
functional tests that call a running REST based WebService.

Most functions require a `t` argument that implements the `forest.T` interface which is a subset of `*testing.T`.

## install

    go get github.com/emicklei/forest

## simple example

    package main

    import (
        "net/http"
        "testing"

        . "github.com/emicklei/forest"
    )

    var github = NewClient("https://api.github.com", new(http.Client))

    func TestForestProjectExists(t *testing.T) {
        cfg := NewConfig("/repos/emicklei/{repo}", "forest").Header("Accept", "application/json")
        r := github.GET(t, cfg)
        ExpectStatus(t, r, 200)
    }

## graphql support

The `GraphQLRequest` can be used to construct a request for a GraphQL endpoint.

```
	query := forest.NewGraphQLRequest(list_matrices_query, "ListMatrices")
	query, err = query.WithVariablesFromString(`
{
	"repositoryID":"99426e24-..........-6bf9770f1fd5",
	"page":{
		"first":20
	},
}`)
	// ... handle error
	cfg := forest.NewRequestConfig(...)
	cfg.Content(query, "application/json")
	r := SkillsAPI.POST(t, cfg)
	ExpectStatus(t, r, 200)
```

- `func NewGraphQLRequest(query, operation string, vars ...Map) GraphQLRequest`
- `func (r GraphQLRequest) WithVariablesFromString(jsonhash string) (GraphQLRequest, error)`

## helper functions

- `func CheckError(t T, err error) bool`
- `func CookieNamed(resp *http.Response, name string) *http.Cookie`
- `func Dump(t T, resp *http.Response)`
- `func Errorf(t *testing.T, format string, args ...interface{})`
- `func ExpectHeader(t T, r *http.Response, name, value string)`
- `func ExpectJSONArray(t T, r *http.Response, callback func(array []interface{}))`
- `func ExpectJSONDocument(t T, r *http.Response, doc interface{})`
- `func ExpectJSONHash(t T, r *http.Response, callback func(hash map[string]interface{}))`
- `func ExpectStatus(t T, r *http.Response, status int) bool`
- `func ExpectString(t T, r *http.Response, callback func(content string))`
- `func ExpectXMLDocument(t T, r *http.Response, doc interface{})`
- `func Fatalf(t *testing.T, format string, args ...interface{})`
- `func JSONArrayPath(t T, r *http.Response, dottedPath string) interface{}`
- `func JSONPath(t T, r *http.Response, dottedPath string) interface{}`
- `func ProcessTemplate(t T, templateContent string, value interface{}) string`
- `func ReadJUnitReport(filename string) (r JUnitReport, err error)`
- `func Scolorf(syntaxCode string, format string, args ...interface{}) string`
- `func SkipUnless(s skippeable, labels ...string)`
- `func VerboseOnFailure(verbose bool)`
- `func XMLPath(t T, r *http.Response, xpath string) interface{}`

## more docs

[Introduction Blog Post](https://ernestmicklei.com/2015/07/testing-your-rest-api-in-go-with-forest/)

© 2016-2025, https://ernestmicklei.com. MIT License. Contributions welcome.
