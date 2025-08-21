package forest

import (
	"net/http"
	"testing"
)

func ExampleJSONArrayPath() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets").Header("Content-Type", "application/json"))
	// if the content looks like this
	// [
	// { "id" : "artreyu", "type" : "tool" }
	// ]
	// then you can verify it using
	if got, want := JSONArrayPath(t, r, ".0.type"), "tool"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleJSONPath() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets/artreyu").Header("Content-Type", "application/json"))
	// if the content looks like this
	// { "id" : "artreyu", "type" : "tool" }
	// then you can verify it using
	if got, want := JSONPath(t, r, ".0.id"), "artreyu"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleExpectJSONArray() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets").Header("Content-Type", "application/json"))
	ExpectJSONArray(t, r, func(array []interface{}) {
		// here you should inspect the array for expected content
		// and use t (*testing.T) to report a failure.
	})
}

func ExampleExpectJSONHash() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets/artreyu/descriptor").Header("Content-Type", "application/json"))
	ExpectJSONHash(t, r, func(hash map[string]interface{}) {
		// here you should inspect the hash for expected content
		// and use t (*testing.T) to report a failure.
	})
}

type YourType struct{}

// How to use the ExpectXMLDocument function on a http response.
func ExampleExpectXMLDocument() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets").Header("Accept", "application/xml"))

	var root YourType // YourType must reflect the expected document structure
	ExpectXMLDocument(t, r, &root)
	// here you should inspect the root (instance of YourType) for expected field values
	// and use t (*testing.T) to report a failure.
}

func ExampleXMLPath() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets/artreyu").Header("Accept", "application/xml"))
	// if the content looks like this
	// <?xml version="1.0" ?>
	// <asset>
	//   <id>artreyu</id>
	//   <type>tool</type>
	// </asset>
	// then you can verify it using
	if got, want := XMLPath(t, r, "/asset/id"), "artreyu"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleExpectStatus() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets/artreyu").Header("Accept", "application/xml"))
	ExpectStatus(t, r, 200)
}

func ExampleExpectHeader() {
	var t *testing.T

	yourAPI := NewClient("http://api.yourservices.com", new(http.Client)) // yourAPI could be a package variable

	r := yourAPI.GET(t, Path("/v1/assets/artreyu").Header("Accept", "application/xml"))
	ExpectHeader(t, r, "Content-Type", "application/xml")
}
