package rest

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

type ConnectionResource struct {
	service application.Logic
}

func (c *ConnectionResource) getFiltered(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	filter := model.ConnectionsFilter{
		Froms:   asFilterParameter(req.QueryParameter("from")),
		Tos:     asFilterParameter(req.QueryParameter("to")),
		Types:   asFilterParameter(req.QueryParameter("type")),
		Centers: asFilterParameter(req.QueryParameter("center"))}
	cons, err := c.service.AllConnections(ctx, filter)
	if err != nil {
		logError("getFilteredConnections", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(cons)
}

func asFilterParameter(param string) (list []string) {
	if len(param) == 0 {
		return list
	}
	return strings.Split(param, ",")
}

func (c *ConnectionResource) putAttribute(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()

	r := bufio.NewReader(req.Request.Body)
	defer req.Request.Body.Close()
	line, err := ioutil.ReadAll(r)
	value := string(line)
	if err != nil {
		logError("putAttribute", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}

	filter := model.ConnectionsFilter{
		Froms: []string{req.PathParameter("from")},
		Tos:   []string{req.PathParameter("to")},
		Types: []string{req.PathParameter("type")}}
	connections, err := c.service.FindAllMatching(ctx, filter)
	if err != nil {
		logError("putAttribute", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	if len(connections) == 0 {
		err = errors.New("no connections found")
		logError("deleteAttribute", err)
		resp.WriteError(http.StatusNotFound, err)
		return
	}
	if len(connections) > 1 {
		err = errors.New("multiple connections found")
		logError("putAttribute", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	connection := &(connections[0]) // TODO use pointer everywhere
	connection.SetAttribute(req.PathParameter("name"), value)
	// now overwrite
	err = c.service.SaveConnection(ctx, *connection, false)
	if err != nil {
		logError("putAttribute", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteAsJson(connection)
}

func (c *ConnectionResource) deleteAttribute(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()

	filter := model.ConnectionsFilter{
		Froms: []string{req.PathParameter("from")},
		Tos:   []string{req.PathParameter("to")},
		Types: []string{req.PathParameter("type")}}
	connections, err := c.service.FindAllMatching(ctx, filter)
	if err != nil {
		logError("deleteAttribute", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	if len(connections) == 0 {
		err = errors.New("no connections found")
		logError("deleteAttribute", err)
		resp.WriteError(http.StatusNotFound, err)
		return
	}
	if len(connections) > 1 {
		err = errors.New("multiple connections found")
		logError("deleteAttribute", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	connection := &(connections[0]) // TODO use pointer everywhere
	connection.DeleteAttribute(req.PathParameter("name"))
	// now overwrite
	err = c.service.SaveConnection(ctx, *connection, false)
	if err != nil {
		logError("deleteAttribute", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteAsJson(connection)
}

func (c *ConnectionResource) put(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	connection := model.Connection{
		From: req.PathParameter("from"),
		To:   req.PathParameter("to"),
		Type: req.PathParameter("type")}
	if err := connection.Validate(); err != nil {
		logError("putConnection", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	err := c.service.SaveConnection(ctx, connection, req.QueryParameter("allowCreate") == "true")
	if err != nil {
		logError("putConnection", err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteHeader(201)
}

func (c *ConnectionResource) delete(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	filter := model.ConnectionsFilter{
		Froms: []string{req.PathParameter("from")},
		Tos:   []string{req.PathParameter("to")},
		Types: []string{req.PathParameter("type")},
	}
	connections, err := c.service.FindAllMatching(ctx, filter)
	if err != nil || len(connections) == 0 {
		logError("deleteConnection", err)
		if err == nil {
			err = errors.New("no connections")
		}
		resp.WriteError(http.StatusNotFound, err)
		return
	}
	for _, each := range connections {
		err = c.service.DeleteConnection(ctx, each)
		if err != nil {
			logError("deleteConnection", err)
			resp.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	resp.WriteHeader(201)
}

func logError(operation string, err error) {
	log.Printf("[landskape-error] %v failed because: %v", operation, err)
}
