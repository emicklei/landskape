package webservice

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"log"
)

type ConnectionService struct {
	restful.WebService
}

func NewConnectionService() *ConnectionService {
	ws := new(ConnectionService)
	ws.Path("/connections").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)
	ws.Route(ws.GET("").To(GetAllConnections))
	return ws
}
func GetAllConnections(req *restful.Request, resp *restful.Response) {
	cons, err := application.SharedLogic.AllConnections()
	if err != nil {
		log.Fatalf("[landskape-error] Request:%v,error:%v", req, err)
		resp.InternalServerError()
	} else {
		resp.WriteEntity(cons)
	}
}
