package main

import (
	"net/http"

	"github.com/emicklei/forest"
)

var api = forest.NewClient("http://localhost:8888", new(http.Client))

type System struct {
	ID    string
	Scope string
}

func main() {
	t := forest.TestingT
	system := System{ID: "sysA"}
	resp := api.POST(t, forest.NewConfig("/{scope}/systems", "test").Content(system, "application/json"))
	forest.Dump(t, resp)

	resp = api.GET(t, forest.NewConfig("/{scope}/systems", "test").Header("Accept", "application/json"))
	forest.Dump(t, resp)
}
