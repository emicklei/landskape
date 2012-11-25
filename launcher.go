package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"log"
	"net/http"
)

func main() {
	restful.Add(application.NewService())
	log.Print(restful.Wadl("http://localhost:9090"))
	log.Fatal(http.ListenAndServe(":9090", nil))
}
