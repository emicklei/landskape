package rest

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

func NewSystemResource(s application.Logic) SystemResource {
	return SystemResource{service: s}
}

func (s SystemResource) Register() {
	ws := new(restful.WebService)
	tags := []string{"systems"}
	ws.Path("/systems").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	idParam := ws.PathParameter("id", "identifier of the system").DataType("string")

	ws.Route(ws.GET("").To(s.getAll).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// docs
		Doc("list all known systems").
		Writes([]model.System{}))

	ws.Route(ws.GET("/{id}").To(s.get).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// docs
		Doc("get the system using its id").
		Param(idParam).
		Writes(model.System{})) // to the response

	ws.Route(ws.PUT("/{id}").To(s.put).
		Param(idParam).
		// docs
		Doc("create the system using its id").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.System{})) // from the request

	ws.Route(ws.DELETE("/{id}").To(s.delete).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// docs
		Doc("delete the system using its id").
		Param(idParam))

	ws.Route(ws.PUT("/{id}/attributes/{name}").To(s.setAttribute).
		Param(idParam).
		Param(ws.PathParameter("name", "name of the attribute. specials = {ui-label,ui-color}")).
		Param(ws.BodyParameter("value", "value of the attribute")).
		// docs
		Doc("set an attribute value").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/{id}/attributes/{name}").To(s.deleteAttribute).
		Param(idParam).
		Param(ws.PathParameter("name", "name of the attribute. specials = {ui-label,ui-color}")).
		// docs
		Doc("delete an attribute value").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	restful.Add(ws)
}
