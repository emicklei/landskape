package webservice

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
	"log"
)

type ApplicationService struct {
	restful.WebService
}

func NewApplicationService() *ApplicationService {
	ws := new(ApplicationService)
	ws.Path("/applications").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)
	ws.Route(ws.GET("").To(GetAllApplications))
	ws.Route(ws.PUT("/{id}").To(CreateApplication))
	return ws
}

func GetAllApplications(req *restful.Request, resp *restful.Response) {
	apps, err := application.SharedLogic.AllApplications()
	if err != nil {
		log.Printf("[landskape-error] AllApplications failed:%v", err)
		resp.WriteError(500, err)
	} else {
		resp.WriteEntity(apps)
	}
}

func CreateApplication(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	app := new(model.Application)
	err := req.ReadEntity(&app)
	if err != nil || app.Id != id {
		log.Printf("[landskape-error] Read failed:%#v", err)
		resp.WriteError(500, err)
	} else {
		_, err = application.SharedLogic.SaveApplication(app)
		if err != nil {
			log.Printf("[landskape-error] Save failed:%#v", err)
			resp.WriteError(500, err)
		}
	}
}
