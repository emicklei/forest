/*Package rat has functions for REST Api testing in Go

This package provides a few simple helper types and functions to create
functional tests that call a running REST based WebService.
A test uses a rat Client that encapsulates a standard http.Client and the base url.
This can be created inside a function, as part of TestMain or as package variable.
Using the Client, you can send http requests and call expectation functions with the response.

Example

		// setup a shared client to your API
		var chatter = rat.NewClient("http://api.chatter.com", new(http.Client))


		func TestGetMessages(t *testing.T) {
			r := chatter.GET(t, rat.Path("/v1/messages").Query("user","zeus"))
			ExpectStatus(t,r,200)
			ExpectJSONArray(t,r,func(messages []interface{}){

				// in the callback you can validate the response structure
				if len(messages) == 0 {
					t.Error("expected messages, got none")
				}
			})
		}

Other expectations

		var root YourDocument
		ExpectXMLDocument(t, r, &root)
		...
		ExpectJSONDocument(t, r, &root)
		...
		ExpectJSONHash(t, r, func(hash map[string]interface{}) {
			...
		})
		...
		ExpectString(t, r, callback func(content string)) {
			...
		})
		...
		ExpectHeader(t, r, name, value string) { ... }

Data access

		value := XMLPath(t, r, "/Root/Child/Value")
		...
		value := JSONPath(t, r, ".Root.Child")
		...
		payload := ProcessTemplate(t, `<Contact><Email>{{.Email}}</Email></Contact>`, contact)


If needed, implement the standard TestMain to do global setup and teardown.

	func TestMain(m *testing.M) {
		// there is no *testing.T available, use an stdout implementation
		t := rat.TestingT

		// setup
		chatter.PUT(t, rat.Path("/v1/messages/{id}",1).Body("<payload>"))
		ExpectStatus(t,r,204)

		exitCode := m.Run()

		// teardown
		chatter.DELETE(t, rat.Path("/v1/messages/{id}",1))
		ExpectStatus(t,r,204)

		os.Exit(exitCode)
	}

Special features

- In contrast to the standard behavior, the Body of a http.Response is made re-readable.
This means one can apply expectations to a response as well as dump the full contents.
- XPath expression support using the [https://godoc.org/launchpad.net/xmlpath] package.
- Colorizes error output (can be configured using package vars).
- Functions can be used in setup and teardown (in body of TestMain).

(c) 2015, http://ernestmicklei.com. MIT License
*/
package rat
