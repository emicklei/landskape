package rest

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

const (
	NO_UPDATE = false
)

type SystemResource struct {
	service application.Logic
}

func (s SystemResource) deleteAttribute(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	name := req.PathParameter("name")
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
	resp.WriteAsJson(papp)
}

func (s SystemResource) setAttribute(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	name := req.PathParameter("name")

	r := bufio.NewReader(req.Request.Body)
	defer req.Request.Body.Close()
	line, err := ioutil.ReadAll(r)
	if err != nil {
		logError("putAttribute", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	value := string(line)

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
	resp.WriteAsJson(papp)
}

// DELETE /systems/{id}
func (s SystemResource) delete(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	err := s.service.DeleteSystem(ctx, id)
	if err != nil {
		if datastore.ErrNoSuchEntity == err {
			resp.WriteError(http.StatusNotFound, err)
		} else {
			resp.WriteError(http.StatusInternalServerError, err)
		}
		return
	}
}

// GET /systems/{id}
func (s SystemResource) get(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	app, err := s.service.GetSystem(ctx, id)
	if err != nil {
		if datastore.ErrNoSuchEntity == err {
			resp.WriteError(http.StatusNotFound, err)
		} else {
			resp.WriteError(http.StatusInternalServerError, err)
		}
		return
	}
	resp.WriteEntity(app)
}

func (s SystemResource) getAll(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	// TODO optimize
	apps, err := s.service.AllSystems(ctx)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	ids := []string{}
	for _, each := range apps {
		ids = append(ids, each.ID)
	}
	resp.WriteEntity(ids)
}

// PUT /systems/{id}
func (s SystemResource) put(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	id := req.PathParameter("id")
	app := model.NewSystem(id)
	err := req.ReadEntity(app)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
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

func (s SystemResource) createAll(req *restful.Request, resp *restful.Response) {
	list := []model.System{}
	if err := req.ReadEntity(&list); err != nil {
		resp.WriteError(400, err)
		return
	}
	ctx := req.Request.Context()
	for _, each := range list {
		each.DBKey = model.NewSystemKey(each.ID)
		// overwrite
		_, err := s.service.SaveSystem(ctx, &each)
		if err != nil {
			log.Println("failed to create system", each.ID)
			resp.WriteError(http.StatusInternalServerError, err)
			return
		}
		log.Println("created system", each.ID)
	}
	resp.WriteHeader(http.StatusCreated)
}
