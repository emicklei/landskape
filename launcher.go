package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/webservice"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

func main() {
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()

	appDao := dao.ApplicationDao{session.DB("landskape").C("applications")}
	conDao := dao.ConnectionDao{session.DB("landskape").C("connections")}
	webservice.SetLogic(application.Logic{appDao, conDao})

	restful.Add(webservice.NewApplicationService())
	restful.Add(webservice.NewConnectionService())
	log.Print(restful.Wadl("http://localhost:9090"))
	log.Fatal(http.ListenAndServe(":9090", nil))
}
