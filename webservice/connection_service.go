package webservice

import (
	"github.com/emicklei/go-restful"
	//	"github.com/emicklei/hopwatch"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
	"log"
	"net/http"
	"strings"
)

func NewConnectionService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/{scope}/connections").
		Param(ws.PathParameter("scope", "organization name to group system and connections")).
		Consumes(restful.MIME_XML).
		Produces(restful.MIME_XML)

	ws.Route(ws.GET("/").
		Doc(`Get all (filtered) connections for all systems and the given scope`).
		Param(ws.QueryParameter("from", "comma separated list of system ids")).
		Param(ws.QueryParameter("to", "comma separated list of system ids")).
		Param(ws.QueryParameter("type", "comma separated list of known connection types")).
		Param(ws.QueryParameter("center", "comma separated list of system ids")).
		To(getFilteredConnections).
		Writes(model.Connection{}))

	ws.Route(ws.PUT("/from/{from}/to/{to}/type/{type}").
		Doc(`Create a new connection using the from,to,type values`).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		Param(ws.QueryParameter("allowCreate", "if true then create any missing systems")).
		To(putConnection).
		Reads(model.Connection{}))

	ws.Route(ws.DELETE("/from/{from}/to/{to}/type/{type}").
		Doc(`Delete an existing connection using the from,to,type values`).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		To(deleteConnection))
	return ws
}

func getFilteredConnections(req *restful.Request, resp *restful.Response) {
	filter := model.ConnectionsFilter{
		Froms:   asFilterParameter(req.QueryParameter("from")),
		Tos:     asFilterParameter(req.QueryParameter("to")),
		Types:   asFilterParameter(req.QueryParameter("type")),
		Centers: asFilterParameter(req.QueryParameter("center"))}
	// hopwatch.Display("filter", filter)
	cons, err := application.SharedLogic.AllConnections(filter)
	if err != nil {
		logError("getFilteredConnections", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(cons)
}

func asFilterParameter(param string) (list []string) {
	if param == "" {
		return list
	}
	return strings.Split(param, ",")
}

func putConnection(req *restful.Request, resp *restful.Response) {
	connection := model.Connection{
		Scope: req.PathParameter("scope"),
		From:  req.PathParameter("from"),
		To:    req.PathParameter("to"),
		Type:  req.PathParameter("type")}
	if err := connection.Validate(); err != nil {
		logError("putConnection", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	err := application.SharedLogic.SaveConnection(connection)
	if err != nil {
		logError("putConnection", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func deleteConnection(req *restful.Request, resp *restful.Response) {
	connection := model.Connection{
		Scope: req.PathParameter("scope"),
		From:  req.PathParameter("from"),
		To:    req.PathParameter("to"),
		Type:  req.PathParameter("type")}
	if err := connection.Validate(); err != nil {
		logError("deleteConnection", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	err := application.SharedLogic.DeleteConnection(connection)
	if err != nil {
		logError("deleteConnection", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func logError(operation string, err error) {
	log.Printf("[landskape-error] %v failed because: %v", operation, err)
}
