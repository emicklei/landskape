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
		Produces(restful.MIME_XML, restful.MIME_JSON).
		Param(ws.PathParameter("scope", "organization name to group system and connections"))

	ws.Route(ws.GET("").To(getAllSystems).
		// docs
		Doc("list all known systems"))
	ws.Route(ws.GET("/{id}").To(getSystem).
		// docs
		Doc("get the system using its id").
		Param(ws.PathParameter("id", "identifier of the system")))
	ws.Route(ws.PUT("/{id}").To(putSystem).
		// docs
		Doc("create the system using its id").
		Param(ws.PathParameter("id", "identifier of the system")))
	ws.Route(ws.POST("").To(postSystem).
		// docs
		Doc("update the system using its id").
		Param(ws.PathParameter("id", "identifier of the system")))
	ws.Route(ws.DELETE("/{id}").To(deleteSystem).
		// docs
		Doc("delete the system using its id").
		Param(ws.PathParameter("id", "identifier of the system")))
	return ws
}

func deleteSystem(req *restful.Request, resp *restful.Response) {
	scope := req.PathParameter("scope")
	id := req.PathParameter("id")
	err := application.SharedLogic.DeleteSystem(scope, id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func getSystem(req *restful.Request, resp *restful.Response) {
	scope := req.PathParameter("scope")
	id := req.PathParameter("id")
	app, err := application.SharedLogic.GetSystem(scope, id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(app)
}

func getAllSystems(req *restful.Request, resp *restful.Response) {
	scope := req.PathParameter("scope")
	apps, err := application.SharedLogic.AllSystems(scope)
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
	if application.SharedLogic.ExistsSystem(scope, id) {
		err := restful.NewError(http.StatusConflict, "System already exists:"+id)
		resp.WriteError(http.StatusConflict, err)
		return
	}
	_, err = application.SharedLogic.SaveSystem(app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}
