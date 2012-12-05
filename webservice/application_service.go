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

type ApplicationService struct {
	restful.WebService
}

func NewApplicationService() *ApplicationService {
	ws := new(ApplicationService)
	ws.Path("/{scope}/applications").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("").To(getAllApplications))
	ws.Route(ws.GET("/{id}").To(getApplication))
	ws.Route(ws.PUT("/{id}").To(putApplication))
	ws.Route(ws.POST("").To(postApplication))
	ws.Route(ws.DELETE("/{id}").To(deleteApplication))
	return ws
}

func deleteApplication(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	err := application.SharedLogic.DeleteApplication(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func getApplication(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	app, err := application.SharedLogic.GetApplication(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(app)
}

func getAllApplications(req *restful.Request, resp *restful.Response) {
	apps, err := application.SharedLogic.AllApplications()
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(apps)
}

func postApplication(req *restful.Request, resp *restful.Response) {
	app := new(model.Application)
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
	_, err = application.SharedLogic.SaveApplication(app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}

func putApplication(req *restful.Request, resp *restful.Response) {
	scope := req.PathParameter("scope")
	id := req.PathParameter("id")
	app := new(model.Application)
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
	if application.SharedLogic.ExistsApplication(id) {
		err := restful.NewError(http.StatusConflict, "Application already exists:"+id)
		resp.WriteError(http.StatusConflict, err)
		return
	}
	_, err = application.SharedLogic.SaveApplication(app)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
	}
}
