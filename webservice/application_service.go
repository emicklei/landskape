package webservice

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"log"
)

// For the complete webservice package
var logic application.Logic

func SetLogic(aLogic application.Logic) {
	logic = aLogic
}

type ApplicationService struct {
	restful.WebService
}

func NewApplicationService() *ApplicationService {
	ws := new(ApplicationService)
	ws.Path("/applications").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)
	ws.Route(ws.GET("").To(GetAllApplications))
	return ws
}
func GetAllApplications(req *restful.Request, resp *restful.Response) {
	apps, err := logic.AllApplications()
	if err != nil {
		log.Fatalf("[landskape-error] Request:%v,error:%v", req, err)
		resp.InternalServerError()
	} else {
		resp.WriteEntity(apps)
	}
}
