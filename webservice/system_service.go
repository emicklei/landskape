package webservice

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
	"net/http"
)

const (
	NO_UPDATE = false
)

func NewSystemService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/{scope}/systems").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("").To(getAllSystems))
	ws.Route(ws.GET("/{id}").To(getSystem))
	ws.Route(ws.PUT("/{id}").To(putSystem))
	ws.Route(ws.POST("").To(postSystem))
	ws.Route(ws.DELETE("/{id}").To(deleteSystem))
	return ws
}

func deleteSystem(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	err := application.SharedLogic.DeleteSystem(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func getSystem(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	app, err := application.SharedLogic.GetSystem(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(app)
}

func getAllSystems(req *restful.Request, resp *restful.Response) {
	apps, err := application.SharedLogic.AllSystems()
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(apps)
}

func postSystem(req *restful.Request, resp *restful.Response) {
	app := new(model.System)
	err := req.ReadEntity(&app)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if app.Id == "" {
		err := restful.NewError(model.MISMATCH_ID, "Id is missing")
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	_, err = application.SharedLogic.SaveSystem(app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}

func putSystem(req *restful.Request, resp *restful.Response) {
	scope := req.PathParameter("scope")
	id := req.PathParameter("id")
	app := new(model.System)
	err := req.ReadEntity(&app)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if app.Id != id {
		err := restful.NewError(model.MISMATCH_ID, fmt.Sprintf("Id mismatch: %v != %v", app.Id, id))
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if app.Scope != "" && app.Scope != scope {
		err := restful.NewError(model.MISMATCH_SCOPE, fmt.Sprintf("Scope mismatch: %v != %v", app.Scope, scope))
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if application.SharedLogic.ExistsSystem(id) {
		err := restful.NewError(http.StatusConflict, "System already exists:"+id)
		resp.WriteError(http.StatusConflict, err)
		return
	}
	_, err = application.SharedLogic.SaveSystem(app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}
