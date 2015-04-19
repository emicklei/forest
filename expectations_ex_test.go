package rat

import "net/http"

func ExampleExpectJSONHash() {
	t := TestingT // t would be a *testing.T

	yourApi := NewClient("http://api.yourservices.com", new(http.Client)) // yourApi could be a package variable

	r := yourApi.GET(t, NewConfig("/v1/assets").Header("Content-Type", "application/json"))
	ExpectJSONHash(t, r, func(hash map[string]interface{}) {
		// here you should inspect the hash for expected content
		// and use t (*testing.T) to report a failure.
	})
}

type YourType struct{}

func ExampleExpectXMLDocument() {
	t := TestingT // t would be a *testing.T

	yourApi := NewClient("http://api.yourservices.com", new(http.Client)) // yourApi could be a package variable

	r := yourApi.GET(t, NewConfig("/v1/assets").Header("Content-Type", "application/xml"))

	var root YourType // YourType must reflect the expected document structure
	ExpectXMLDocument(t, r, &root)
	// here you should inspect the root (instance of YourType) for expected field values
	// and use t (*testing.T) to report a failure.
}
