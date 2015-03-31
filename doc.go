/*
Package rat, rest api testing in Go

This package provides a few simple helper types and functions to create
functional tests that call a running REST based WebService.
A test uses a rat Client that encapsulates a standard http.Client and the base url.
This can be created inside a function, as part of TestMain or as package variable.
Using the Client, you can send http requests and call expectation functions with the response.

Example

	// setup a shared client to your API
	var chatter = rat.NewClient("http://api.chatter.com", new(http.Client))


	func TestGetMessages(t *testing.T) {
		r := chatter.GET(t, rat.NewConfig("/v1/messages?user=zeus"))
		chatter.ExpectStatus(t,r,200)
		chatter.ExpectJsonArray(t,r,func(messages []interface{}){

			// in the callback you can validate the response structure
			if len(messages) == 0 {
				t.Error("expected messages, got none")
			}
		})
	}

If needed, implement the standard TestMain to do global setup and teardown.

	func TestMain(m *testing.M) {
		// there is no *testing.T available, use the failing one
		t := rat.FailingT

		// setup
		chatter.PUT(t, rat.NewConfig("/v1/messages/1").Body("<payload>"))
		chatter.ExpectStatus(t,r,204)

		exitCode := m.Run()

		// teardown
		chatter.DELETE(t, rat.NewConfig("/v1/messages/1"))
		chatter.ExpectStatus(t,r,204)

		os.Exit(exitCode)
	}


(c) 2015, http://ernestmicklei.com. MIT License
*/
package rat
