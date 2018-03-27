# Landskape - maintain and visualize connected systems 

### Installation

		go get -u github.com/emicklei/landskape

### Run locally

    dev_appserver.py app.yaml

Swagger

	http://localhost:8888/swagger-ui/?url=http://localhost:8888/api-docs.json

### Build

To build the project locally and test it.

	go get -u github.com/jteeuwen/go-bindata/...

Make sure $GOPATH/bin is on your $PATH.

	go generate

(c) 2012-2017, ernestmicklei.com. MIT License