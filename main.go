package main

import (
	"context"
	"flag"
	"fmt"
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
	log.Print("[landskape] startup...")
	flag.Parse()
	props, _ := properties.Load(*propertiesFile)
	log.Println("[landskape]", props)
	log.Println("[landskape] GOOGLE_CLOUD_PROJECT=", os.Getenv("GOOGLE_CLOUD_PROJECT"))
	initSelfdiagnose()

	// prepare datastore
	ds, err := datastore.NewClient(context.Background(), os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal("datastore client creation failed, missing GOOGLE_CLOUD_PROJECT ?", err)
	}

	appDao := dao.NewSystemDao(ds)
	conDao := dao.NewConnectionDao(ds)
	service := application.Logic{SystemDao: appDao, ConnectionDao: conDao}

	wsSystem := rest.NewSystemResource(service).NewWebService()
	wsConnection := rest.NewConnectionResource(service).NewWebService()
	wsDiagram := rest.NewDiagramService(service)

	wsSystem.Filter(apiKeyAuthenticate)
	wsConnection.Filter(apiKeyAuthenticate)
	wsDiagram.Filter(apiKeyAuthenticate)

	restful.Add(wsSystem)
	restful.Add(wsConnection)
	restful.Add(wsDiagram)

	// for graphical diagrams
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
	restful.DefaultContainer.Add(IndexService())

	port := props["http.server.port"]
	publicHost := "localhost"
	log.Println(fmt.Sprintf("[landskape] swagger http://%s:%s/swagger-ui/?url=http://%s:%s/api-docs.json", publicHost, port, publicHost, port))
	log.Printf("[landskape] ready to serve on %v\n", basePath)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Landskape",
			Description: "Logical communication diagrams of system infrastructure",
			Contact: &spec.ContactInfo{
				Name: "PhilemonWorks",
				URL:  "https://github.com/emicklei/landskape",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	// setup security definitions
	swo.SecurityDefinitions = map[string]*spec.SecurityScheme{
		"api_key": spec.APIKeyAuth("api_key", "query"),
	}
	swo.Security = append(swo.Security, map[string][]string{
		"api_key": []string{},
	})
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "systems",
		Description: "Managing Systems"}}, spec.Tag{TagProps: spec.TagProps{
		Name:        "connections",
		Description: "Managing Connections"}}, spec.Tag{TagProps: spec.TagProps{
		Name:        "diagrams",
		Description: "Display graphs"}}}
}
