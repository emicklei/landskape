package rest

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

const (
	NO_UPDATE = false
)

type SystemResource struct {
	service application.Logic
}

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
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// docs
		Doc("create the system using its id").
		Param(idParam).
		Reads(model.System{})) // from the request

	ws.Route(ws.POST("").To(s.post).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// docs
		Doc("update the system using its id").
		Param(idParam).
		Reads(model.System{})) // from the request

	ws.Route(ws.DELETE("/{id}").To(s.delete).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// docs
		Doc("delete the system using its id").
		Param(idParam))

	ws.Route(ws.PUT("/{id}/attributes").To(s.setAttribute).
		Param(idParam).
		Param(ws.QueryParameter("name", "name of the attribute. specials = {ui-label,ui-color}")).
		Param(ws.QueryParameter("value", "value of the attribute")).
		// docs
		Doc("set an attribute value").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/{id}/attributes").To(s.deleteAttribute).
		Param(idParam).
		Param(ws.QueryParameter("name", "name of the attribute. specials = {ui-label,ui-color}")).
		// docs
		Doc("set an attribute value").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	restful.Add(ws)
}

func (s SystemResource) deleteAttribute(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	name := req.QueryParameter("name")
	if len(name) == 0 {
		resp.WriteErrorString(400, "name cannot be empty")
		return
	}
	app, err := s.service.GetSystem(ctx, id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	papp := &app // TODO
	papp.DeleteAttribute(name)
	_, err = s.service.SaveSystem(ctx, papp)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}

func (s SystemResource) setAttribute(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	name := req.QueryParameter("name")
	if len(name) == 0 {
		resp.WriteErrorString(400, "name cannot be empty")
		return
	}
	value := req.QueryParameter("value")
	app, err := s.service.GetSystem(ctx, id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	papp := &app // TODO
	papp.SetAttribute(name, value)
	_, err = s.service.SaveSystem(ctx, papp)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
	resp.WriteEntity(papp)
}

// DELETE /systems/{id}
func (s SystemResource) delete(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	err := s.service.DeleteSystem(ctx, id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

// GET /systems/{id}
func (s SystemResource) get(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	app, err := s.service.GetSystem(ctx, id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(app)
}

func (s SystemResource) getAll(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	apps, err := s.service.AllSystems(ctx)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(apps)
}

// POST /systems/
func (s SystemResource) post(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	app := new(model.System)
	err := req.ReadEntity(app)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if len(app.ID) == 0 {
		err := restful.NewError(model.MISMATCH_ID, "Id is missing")
		resp.WriteServiceError(http.StatusBadRequest, err)
		return
	}
	_, err = s.service.SaveSystem(ctx, app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT /systems/{id}
func (s SystemResource) put(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	app := new(model.System)
	err := req.ReadEntity(app)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if app.ID != id {
		err := restful.NewError(model.MISMATCH_ID, fmt.Sprintf("Id mismatch: %v != %v", app.ID, id))
		resp.WriteServiceError(http.StatusBadRequest, err)
		return
	}
	if s.service.ExistsSystem(ctx, id) {
		err := restful.NewError(http.StatusConflict, "System already exists:"+id)
		resp.WriteServiceError(http.StatusConflict, err)
		return
	}
	_, err = s.service.SaveSystem(ctx, app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteHeader(http.StatusCreated)
}
