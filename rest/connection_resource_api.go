package rest

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

func NewConnectionResource(s application.Logic) ConnectionResource {
	return ConnectionResource{service: s}
}

func (c ConnectionResource) Register() {
	ws := new(restful.WebService)
	tags := []string{"connections"}

	ws.Path("/connections").
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").
		Doc(`Get all (filtered) connections for all systems and the given scope`).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("from", "comma separated list of system ids")).
		Param(ws.QueryParameter("to", "comma separated list of system ids")).
		Param(ws.QueryParameter("type", "comma separated list of known connection types")).
		Param(ws.QueryParameter("center", "comma separated list of system ids")).
		To(c.getFiltered).
		Writes([]model.Connection{}))

	ws.Route(ws.PUT("/from/{from}/to/{to}/type/{type}").
		Doc(`Create a new connection using the from,to,type values`).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		Param(ws.QueryParameter("allowCreate", "if true then create any missing systems")).
		To(c.put))

	ws.Route(ws.PUT("/from/{from}/to/{to}/type/{type}/attributes/{name}").
		Doc(`Create a new connection attribute using the from,to,type values`).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		Param(ws.PathParameter("name", "name of the attribute. specials = {ui-label,ui-color}")).
		Param(ws.BodyParameter("value", "value of the attribute")).
		To(c.putAttribute))

	ws.Route(ws.DELETE("/from/{from}/to/{to}/type/{type}/attributes/{name}").
		Doc(`Delete a new connection attribute using the from,to,type values`).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		Param(ws.PathParameter("name", "name of the attribute. specials = {ui-label,ui-color}")).
		To(c.deleteAttribute))

	ws.Route(ws.DELETE("/from/{from}/to/{to}/type/{type}").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc(`Delete an existing connection using the from,to,type values`).
		Param(ws.PathParameter("from", "system id")).
		Param(ws.PathParameter("to", "system id")).
		Param(ws.PathParameter("type", "indicate type of connection, e.g. http,jdbc,ftp,aq")).
		To(c.delete))

	restful.Add(ws)
}
