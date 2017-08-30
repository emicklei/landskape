package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/dmotylev/goproperties"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/dao/tiedot"
	"github.com/emicklei/landskape/webservice"
	"github.com/go-openapi/spec"
)

var propertiesFile = flag.String("config", "landskape.properties", "the configuration file")

func main() {
	log.Print("[landskape] service startup...")
	flag.Parse()
	props, _ := properties.Load(*propertiesFile)

	tdb, err := db.OpenDB("./tiedot.db")
	if err != nil {
		log.Fatal("opendb failed", err)
	}
	if err := tdb.Create("systems"); err != nil {
		log.Println("create systems failed ", err)
	}
	if err := tdb.Create("connections"); err != nil {
		log.Println("create connections failed ", err)
	}

	appDao := tiedot.SystemDao{tdb.Use("systems")}
	conDao := tiedot.ConnectionDao{tdb.Use("connections")}
	application.SharedLogic = application.Logic{appDao, conDao}

	webservice.SystemResource{application.SharedLogic}.Register()
	webservice.ConnectionResource{application.SharedLogic}.Register()

	// graphical diagrams
	restful.Add(webservice.NewDiagramService())
	webservice.DotConfig["binpath"] = props["dot.path"]
	webservice.DotConfig["tmp"] = props["dot.tmp"]

	// expose api using swagger
	basePath := "http://" + props["http.server.host"] + ":" + props["http.server.port"]

	config := restfulspec.Config{
		WebServices:    restful.RegisteredWebServices(),
		WebServicesURL: fmt.Sprintf("%s%s", basePath, props["swagger.path"]),
		APIPath:        props["swagger.api"],
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	http.Handle("/doc/", http.StripPrefix("/doc/", http.FileServer(http.Dir("/Users/emicklei/xProjects/swagger-ui/dist"))))

	log.Printf("[landskape] ready to serve on %v\n", basePath)
	log.Fatal(http.ListenAndServe(":"+props["http.server.port"], nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Landskape",
			Description: "Flow diagrams for infrastructure",
			Contact: &spec.ContactInfo{
				Name:  "john",
				Email: "john@doe.rp",
				URL:   "http://johndoe.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "systems",
		Description: "Managing Systems"}}, spec.Tag{TagProps: spec.TagProps{
		Name:        "connections",
		Description: "Managing Connections"}}}
}
