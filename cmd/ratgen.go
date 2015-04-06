package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/emicklei/go-restful/swagger"
)

var swaggerJsonUrl = flag.String("url", "http://localhost:8080/apidocs.json", "full URL of the Swagger JSON listing")
var targetDirectory = flag.String("o", "/tmp", "directory to generate test files in")

func main() {
	flag.Parse()
	listing := getListing()

	// setup
	where := path.Join(*targetDirectory, "setup.go")
	s, err := os.Create(where)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	setup.Execute(s, basePathFrom(*swaggerJsonUrl))
	log.Printf("[ratgen] written %s\n", where)

	for _, each := range listing.Apis {
		decl := getApiDeclaration(each.Path)

		// tes file per operation
		for _, api := range decl.Apis {
			for _, op := range api.Operations {
				writeTest(api, op)
			}
		}
	}
}

type testParams struct {
	Name, Method, Path string
	Status             int
}

func writeTest(api swagger.Api, op swagger.Operation) {
	data := testParams{
		Name:   op.Nickname,
		Method: op.Method,
		Path:   api.Path,
		Status: status(op),
	}
	context := sanitize(api.Path)
	where := path.Join(*targetDirectory, "test_"+context+"_"+op.Nickname+"_test.go")
	o, err := os.Create(where)
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()
	if hasPathParams(api.Path) {
		fun_pathparams.Execute(o, data)
	} else {
		fun.Execute(o, data)
	}
	log.Printf("[ratgen] written %s\n", where)
}

func basePathFrom(url string) string {
	// guess
	j := strings.Index(url, "/apidocs.json")
	return url[:j]
}

func sanitize(resourcepath string) string {
	withoutSlashes := strings.Replace(resourcepath, "/", "", -1)
	curly := strings.Index(withoutSlashes, "{")
	if curly > -1 {
		return withoutSlashes[:curly]
	}
	return withoutSlashes
}

func hasPathParams(resourcepath string) bool {
	return strings.Index(resourcepath, "{") != -1
}

func status(op swagger.Operation) int {
	if len(op.ResponseMessages) > 0 {
		return op.ResponseMessages[0].Code

	}
	m := op.Method
	if m == "PUT" || m == "DELETE" {
		return 204
	}
	return 200
}
