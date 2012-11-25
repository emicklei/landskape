package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/webservice"
	"log"
	"net/http"
)

func main() {
	webservice.SetLogic(application.Logic{})

	restful.Add(webservice.NewApplicationService())
	restful.Add(webservice.NewConnectionService())
	log.Print(restful.Wadl("http://localhost:9090"))
	log.Fatal(http.ListenAndServe(":9090", nil))
}
