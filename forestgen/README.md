# forestgen - test file generator

It requires a Swagger 1.2 API JSON endpoint.

	forestgen -url "http://localhost:8080/apidocs.json" -o "/tmp"
	
### Installation

	cd  $GOPATH/src/github.com/emicklei/forest/forestgen
	go build -o $GOPATH/bin/forestgen
	chmod +x $GOPATH/bin/forestgen
	
Make sure you have `$GOPATH/bin` on your PATH.
Now you can verify the availability of the tool.
	
	forestgen -help
	
(c) 2015, http://ernestmicklei.com. MIT License	