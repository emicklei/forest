# rat - rest api testing in Go

	// setup a shared client to your API
	var chatter = rat.NewClient("http://api.chatter.com", new(http.Client))
	
	// if needed, implement the standard TestMain to do global setup and teardown
	func TestMain(m *testing.M) {
		// there is no *testing.T available, use the logging one
		t := rat.LoggingT	
		
		// setup
		chatter.PUT(t, "/v1/messages/1", rat.NewRequestConfig().Body("<payload>"))
		chatter.ExpectStatus(t,r,204)
		
		exitCode := m.Run()
		
		// teardown
		chatter.DELETE(t,"/v1/messages/1")
		chatter.ExpectStatus(t,r,204)
		
		os.Exit(exitCode)
	}
	
	func TestGetMessages(t *testing.T) {
		r := chatter.GET(t,"/v1/messages?user=zeus")	
		chatter.ExpectStatus(t,r,200)
		chatter.ExpectJsonArray(t,r,func(messages []interface{}){
			
			// in the callback you can validate the response structure
			if len(messages) == 0 {
				t.Error("expected messages, got none")
			}
		})
	}
		
(c) 2015, http://ernestmicklei.com. MIT License	