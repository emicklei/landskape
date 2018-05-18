# Landskape - explore connected systems 

## Installation

	go get -u github.com/emicklei/landskape

### Run locally

    make run


## Attributes

Both Systems and Connections can have arbitrary attributes.
Some attributes cause special behavior when creating visual diagrams.
Note that attribute names are case-sensitive. This is usually true for attribute values as well, unless noted.

### cluster

Setting the `cluster` attribute on a system allows you to create diagrams with clustering based on the value of this attribute.
For example, after adding `team = A` to System A and `team = B` to System B then creating a diagram with `cluster=team` will group all Systems based on the `team` value.

### graphs

https://www.graphviz.org/doc/info/attrs.html

### backup & restore

You can use this `Makefile` to get data out and in landskape using the REST API.

	backup: get_sys get_con

	restore: put_sys put_con

	get_sys:
		curl http://localhost:8080/v1/systems > systems.json

	put_sys:
		curl -X POST -H "Content-Type:application/json" http://localhost:8080/v1/systems -d @systems.json

	get_con:
		curl http://localhost:8080/v1/connections > connections.json

	put_con:
		curl -X POST -H "Content-Type:application/json" http://localhost:8080/v1/connections -d @connections.json
		
## Deloy to GCP

- enable Google App Engine Flexible Environment 		

## Development

To build the project locally and test it.

	go get -u github.com/jteeuwen/go-bindata/...

Make sure $GOPATH/bin is on your $PATH.

	go generate
	
See `Makefile` for local docker instructions.


(c) 2012-2018, ernestmicklei.com. MIT License