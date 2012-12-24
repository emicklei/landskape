package webservice

import (
	"github.com/emicklei/go-restful"

//	"github.com/emicklei/landskape/dao"
//	"github.com/emicklei/landskape/model"
)

func NewDiagramService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/{scope}/diagram").
		Param(ws.PathParameter("scope", "organization name to group system and connections")).
		Produces("text/plain")
	ws.Route(ws.GET("/").To(computeDiagram))
	return ws
}

func computeDiagram(req *restful.Request, resp *restful.Response) {
	//filter := model.ConnectionsFilter{}
	//connections := application.SharedLogic.AllConnections(filter)
	//dotBuilder := 
}
