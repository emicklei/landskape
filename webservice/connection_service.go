package webservice

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
	"log"
	"net/http"
)

type ConnectionService struct {
	restful.WebService
}

func NewConnectionService() *ConnectionService {
	ws := new(ConnectionService)
	ws.Path("/{scope}/connections").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)
	ws.Route(ws.GET("").
		To(GetFilteredConnections).
		Doc(`Get all (filtered) connections for all applications and the given scope`))
	ws.Route(ws.PUT("/from/{from}/to/{to}/type/{type}").
		To(PutConnection).
		Doc(`Create a new connection using the from,to,type values`))
	return ws
}
func GetFilteredConnections(req *restful.Request, resp *restful.Response) {
	filter := model.ConnectionsFilter{
		From:   req.QueryParameter("From"),
		To:     req.QueryParameter("To"),
		Type:   req.QueryParameter("Type"),
		Center: req.QueryParameter("Center")}
	log.Printf("filter:%#v", filter)
	cons, err := application.SharedLogic.AllConnections(filter)
	if err != nil {
		logError("GetFilteredConnections", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(cons)
}

func PutConnection(req *restful.Request, resp *restful.Response) {
	connection := model.Connection{
		Scope: req.PathParameter("scope"),
		From:  req.PathParameter("from"),
		To:    req.PathParameter("to"),
		Type:  req.PathParameter("type")}
	if err := connection.Validate(); err != nil {
		logError("PutConnection", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	err := application.SharedLogic.SaveConnection(connection)
	if err != nil {
		logError("PutConnection", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func logError(operation string, err error) {
	log.Printf("[landskape-error] %v failed because: %v", operation, err)
}
