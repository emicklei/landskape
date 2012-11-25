package webservice

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

type ConnectionService struct {
	restful.WebService
	logic application.Logic
}

func NewConnectionService() *ConnectionService {
	ws := new(ConnectionService)
	ws.Path("/connections").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)
	ws.Route(ws.GET("").To(GetAllConnections))
	return ws
}
func GetAllConnections(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(model.Connection{})
}
