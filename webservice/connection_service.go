package webservice

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
	"log"
	"net/http"
	"strings"
)

func NewConnectionService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/{scope}/connections").
		PathParam("scope", "organization name to group applications and connections").
		Consumes(restful.MIME_XML).
		Produces(restful.MIME_XML)

	ws.Route(ws.GET("?from={from}&to={to}&type={type}&center={center}").
		Doc(`Get all (filtered) connections for all applications and the given scope`).
		QueryParam("from", "comma separated list of application ids").
		QueryParam("to", "comma separated list of application ids").
		QueryParam("type", "comma separated list of known connection types").
		QueryParam("center", "comma separated list of application ids").
		To(getFilteredConnections).
		Writes(model.Connection{}))

	ws.Route(ws.PUT("/from/{from}/to/{to}/type/{type}?allowCreate={true|false}").
		Doc(`Create a new connection using the from,to,type values`).
		PathParam("from", "application id").
		PathParam("to", "application id").
		PathParam("type", "indicate type of connection, e.g. http,jdbc,ftp,aq").
		QueryParam("allowCreate", "if true then create any missing applications").
		To(putConnection).
		Reads(model.Connection{}))

	ws.Route(ws.DELETE("/from/{from}/to/{to}/type/{type}").
		Doc(`Delete an existing connection using the from,to,type values`).
		PathParam("from", "application id").
		PathParam("to", "application id").
		PathParam("type", "indicate type of connection, e.g. http,jdbc,ftp,aq").
		To(deleteConnection))
	return ws
}

func getFilteredConnections(req *restful.Request, resp *restful.Response) {
	filter := model.ConnectionsFilter{
		Froms:   strings.Split(req.QueryParameter("from"), ","),
		Tos:     strings.Split(req.QueryParameter("to"), ","),
		Types:   strings.Split(req.QueryParameter("type"), ","),
		Centers: strings.Split(req.QueryParameter("center"), ",")}
	log.Printf("filter:%#v", filter)
	cons, err := application.SharedLogic.AllConnections(filter)
	if err != nil {
		logError("getFilteredConnections", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(cons)
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
