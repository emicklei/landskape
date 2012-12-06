package main

import (
	"flag"
	"github.com/dmotylev/goproperties/src/goproperties"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/webservice"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"os"
)

var propertiesFile = flag.String("config", "landskape.properties", "the configuration file")

func main() {
	flag.Parse()
	props, _ := readProperties(*propertiesFile)
	session, _ := mgo.Dial(props["mongo.connection"])
	defer session.Close()

	appDao := dao.ApplicationDao{session.DB(props["mongo.database"]).C("applications")}
	conDao := dao.ConnectionDao{session.DB(props["mongo.database"]).C("connections")}
	application.SharedLogic = application.Logic{appDao, conDao}

	restful.Add(webservice.NewApplicationService())
	restful.Add(webservice.NewConnectionService())

	// expose api using swagger
	basePath := "http://" + props["http.server.host"] + ":" + props["http.server.port"]
	restful.Add(restful.NewSwaggerService(basePath, "/api-docs.json"))

	log.Fatal(http.ListenAndServe(":"+props["http.server.port"], nil))
}

func readProperties(filename string) (map[string]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open %s: %s", filename, err)
	}
	defer f.Close()
	return goproperties.Load(f)
}
