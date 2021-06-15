# forest - for testing REST api-s in Go

[![Build Status](https://travis-ci.org/emicklei/forest.png)](https://travis-ci.org/emicklei/forest)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/forest)](https://goreportcard.com/report/github.com/emicklei/forest)
[![GoDoc](https://godoc.org/github.com/emicklei/forest?status.svg)](https://godoc.org/github.com/emicklei/forest)

This package provides a few simple helper types and functions to create
functional tests that call a running REST based WebService.

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

    TOWRITE

## other helper functions

    func ExpectHeader(t T, r *http.Response, name, value string)
    func ExpectJSONArray(t T, r *http.Response, callback func(array []interface{}))
    func ExpectJSONDocument(t T, r *http.Response, doc interface{})
    func ExpectJSONHash(t T, r *http.Response, callback func(hash map[string]interface{}))
    func ExpectStatus(t T, r *http.Response, status int) bool
    func ExpectString(t T, r *http.Response, callback func(content string))
    func ExpectXMLDocument(t T, r *http.Response, doc interface{})
    func JSONArrayPath(t T, r *http.Response, dottedPath string) interface{}
    func JSONPath(t T, r *http.Response, dottedPath string) interface{}
    func ProcessTemplate(t T, templateContent string, value interface{}) string
    func Scolorf(syntaxCode string, format string, args ...interface{}) string
    func SkipUnless(s skippeable, labels ...string)
    func XMLPath(t T, r *http.Response, xpath string) interface{}
    func Dump(t T, resp *http.Response)

## more docs

[Introduction Blog Post](http://ernestmicklei.com/2015/07/testing-your-rest-api-in-go-with-forest/)
		
© 2016+, http://ernestmicklei.com. MIT License. Contributions welcome.	 
