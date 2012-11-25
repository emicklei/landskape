package application

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/model"
)

type ApplicationService struct {
	restful.WebService
}

func NewService() *ApplicationService {
	ws := new(ApplicationService)
	ws.Path("/applications").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)
	ws.Route(ws.GET("").To(GetAllApplications))
	return ws
}
func GetAllApplications(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(model.Application{})
}