package rest

import (
	"log"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

type ConnectionResource struct {
	service application.Logic
}

func NewConnectionResource(s application.Logic) ConnectionResource {
	return ConnectionResource{service: s}
}

func (c ConnectionResource) Register() {
	ws := new(restful.WebService)
	tags := []string{"connections"}

	ws.Path("/{scope}/connections").
		Param(ws.PathParameter("scope", "organization name to group system and connections")).
		Consumes(restful.MIME_XML).
		Produces(restful.MIME_XML)

	ws.Route(ws.GET("/").
		Doc(`Get all (filtered) connections for all systems and the given scope`).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("from", "comma separated list of system ids")).
		Param(ws.QueryParameter("to", "comma separated list of system ids")).
		Param(ws.QueryParameter("type", "comma separated list of known connection types")).
		Param(ws.QueryParameter("center", "comma separated list of system ids")).
		To(c.getFiltered).
		Writes(model.Connection{}))

	ws.Route(ws.PUT("/from/{from}/to/{to}/type/{type}").
		Doc(`Create a new connection using the from,to,type values`).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		Param(ws.QueryParameter("allowCreate", "if true then create any missing systems")).
		To(c.put).
		Reads(model.Connection{}))

	ws.Route(ws.DELETE("/from/{from}/to/{to}/type/{type}").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc(`Delete an existing connection using the from,to,type values`).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		To(c.delete))

	restful.Add(ws)
}

func (c *ConnectionResource) getFiltered(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	scope := req.PathParameter("scope")
	filter := model.ConnectionsFilter{
		Froms:   asFilterParameter(req.QueryParameter("from")),
		Tos:     asFilterParameter(req.QueryParameter("to")),
		Types:   asFilterParameter(req.QueryParameter("type")),
		Centers: asFilterParameter(req.QueryParameter("center"))}
	// hopwatch.Display("filter", filter)
	cons, err := c.service.AllConnections(ctx, scope, filter)
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

func (c *ConnectionResource) put(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
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
	err := c.service.SaveConnection(ctx, connection)
	if err != nil {
		logError("putConnection", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func (c *ConnectionResource) delete(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
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
	err := c.service.DeleteConnection(ctx, connection)
	if err != nil {
		logError("deleteConnection", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func logError(operation string, err error) {
	log.Printf("[landskape-error] %v failed because: %v", operation, err)
}
