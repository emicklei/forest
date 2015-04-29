package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/swagger"
)

func getListing() *swagger.ResourceListing {
	r, err := http.Get(*swaggerJsonUrl)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf("Getting swagger listing failed: unable to read response body:%v", err)
	}
	var listing swagger.ResourceListing
	err = json.Unmarshal(data, &listing)
	if err != nil {
		log.Fatalf("Parsing swagger listing failed:%v", err)
	}
	return &listing
}

func getApiDeclaration(path string) *swagger.ApiDeclaration {
	r, err := http.Get(*swaggerJsonUrl + path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf("Getting swagger api failed: unable to read response body:%v", err)
	}
	var api swagger.ApiDeclaration
	err = json.Unmarshal(data, &api)
	if err != nil {
		log.Fatalf("Parsing swagger api failed:%v", err)
	}
	return &api
}
