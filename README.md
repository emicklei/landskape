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


## Attributes

Both Systems and Connections can have arbitrary attributes.
Some attributes cause special behavior when creating visual diagrams.
Note that attribute names are case-sensitive. This is usually true for attribute values as well, unless noted.

### cluster

Setting the `cluster` attribute on a system allows you to create diagrams with clustering based on the value of this attribute.
For example, after adding `team = A` to System A and `team = B` to System B then creating a diagram with `cluster=team` will group all Systems based on the `team` value.

### graphs

https://www.graphviz.org/doc/info/attrs.html

(c) 2012-2017, ernestmicklei.com. MIT License