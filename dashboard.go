package main

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func indexPage(req *restful.Request, resp *restful.Response) {
	r := req.Request
	s := "http"
	if proto := req.HeaderParameter("X-Forwarded-Proto"); len(proto) > 0 {
		s = proto
	}
	http.Redirect(resp.ResponseWriter, req.Request, fmt.Sprintf("/swagger-ui/?url=%s://%s/api-docs.json", s, r.Host), 302)
}

func IndexService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(indexPage))
	return ws
}
