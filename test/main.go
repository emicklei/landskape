package main

import (
	"encoding/json"
	"net/http"

	"github.com/emicklei/forest"
)

// go run main.go

var api = forest.NewClient("http://localhost:8888", new(http.Client))

type System struct {
	ID string
}
type Connection struct {
	From, To, Type string
}

func main() {
	flushConnections()
}

func createAll() {
	t := forest.TestingT
	system := System{ID: "sysA"}
	resp := api.PUT(t, forest.NewConfig("/systems/{id}", system.ID).Content(system, "application/json"))
	forest.Dump(t, resp)

	{
		system := System{ID: "sysB"}
		resp := api.PUT(t, forest.NewConfig("/systems/{id}", system.ID).Content(system, "application/json"))
		forest.Dump(t, resp)
	}

	{
		resp := api.PUT(t, forest.NewConfig("/connections/from/{from}/to/{to}/type/{type}",
			"sysA",
			"sysB",
			"jdbc"))
		forest.Dump(t, resp)
	}

	resp = api.GET(t, systemConfig())
	forest.Dump(t, resp)
}

func systemConfig() *forest.RequestConfig {
	return forest.NewConfig("/systems").Header("Accept", "application/json")
}

func flush() {
	flushConnections()
	//flushSystems()
}
func flushConnections() {
	t := forest.TestingT
	resp := api.GET(t, forest.NewConfig("/connections").Header("Accept", "application/json"))
	list := []Connection{}
	json.NewDecoder(resp.Body).Decode(&list)
	for _, each := range list {
		api.DELETE(t, forest.NewConfig("/connections/from/{from}/to/{to}/type/{type}", each.From, each.To, each.Type))
	}
}
