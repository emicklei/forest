# ratgen - test file generator

It requires a Swagger 1.2 API JSON endpoint.

	ratgen -url "http://localhost:8080/apidocs.json" -o "/tmp"
	
### Installation

	cd  $GOPATH/src/github.com/emicklei/rat/ratgen
	go build -o $GOPATH/bin/ratgen
	chmod +x $GOPATH/bin/ratgen
	
Make sure you have `$GOPATH/bin` on your PATH.
Now you can verify the availability of the tool.
	
	ratgen -help
	
(c) 2015, http://ernestmicklei.com. MIT License	