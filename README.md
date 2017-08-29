#Landskape

###Webservice to maintain and visualize connected systems

# REWORK in progress !!!

Requirements
	  
	MongoDB db version v2.2.0+     
	Swagger-UI 1.1.7+

Installation

		go get -u github.com/emicklei/landskape

Configuration example

	mongo.connection=localhost:27017
	mongo.database=landskape
                     
	http.server.host=localhost
	http.server.port=9090     
	                         
	swagger.api=/apidocs.json
    swagger.ui=/apidocs/
	swagger.folder=/Users/emicklei/Downloads/swagger-ui-1.1.7

Flags

		-config  the configuration file
	
Start

	go run main.go


![landskape api (swagger)](https://s3.amazonaws.com/public.philemonworks.com/landskape-api-swagger.png)

(c) 2012-2017, ernestmicklei.com. MIT License