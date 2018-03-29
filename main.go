package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/dmotylev/goproperties"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/rest"
	"github.com/go-openapi/spec"
)

//go:generate go-bindata -pkg main swagger-ui/...

var propertiesFile = flag.String("config", "landskape.properties", "the configuration file")

func main() {
	log.Print("[landskape] service startup...")
	flag.Parse()
	props, _ := properties.Load(*propertiesFile)

	// prepare datastore
	ds, err := datastore.NewClient(context.Background(), os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal("datastore client creation failed", err)
	}

	appDao := dao.NewSystemDao(ds)
	conDao := dao.NewConnectionDao(ds)
	service := application.Logic{appDao, conDao}

	rest.NewSystemResource(service).Register()
	rest.NewConnectionResource(service).Register()

	// graphical diagrams
	restful.Add(rest.NewDiagramService(service))
	rest.DotConfig["binpath"] = props["dot.path"]
	rest.DotConfig["tmp"] = props["dot.tmp"]

	// expose api using swagger
	basePath := "http://" + props["http.server.host"] + ":" + props["http.server.port"]

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     props["swagger.api"],
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// static file serving
	swaggerUI := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "swagger-ui/dist"}
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(swaggerUI)))

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
