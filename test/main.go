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
	resp := api.POST(t, forest.NewConfig("/systems").Content(system, "application/json"))
	forest.Dump(t, resp)

	{
		system := System{ID: "sysB"}
		resp := api.POST(t, forest.NewConfig("/systems").Content(system, "application/json"))
		forest.Dump(t, resp)
	}

	{
		resp := api.PUT(t, forest.NewConfig("/connections/from/{from}/to/{to}/type/{type}",
			"sysA",
			"sysB",
			"jdbc"))
		forest.Dump(t, resp)
	}

	resp = api.GET(t, forest.NewConfig("/systems").Header("Accept", "application/json"))
	forest.Dump(t, resp)
}
