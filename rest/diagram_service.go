package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/landskape/application"
	"github.com/emicklei/landskape/model"
)

var DotConfig = map[string]string{}

type DiagramResource struct {
	service application.Logic
}

func (d DiagramResource) computeDiagram(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	filter := model.ConnectionsFilter{
		Froms:   asFilterParameter(req.QueryParameter("from")),
		Tos:     asFilterParameter(req.QueryParameter("to")),
		Types:   asFilterParameter(req.QueryParameter("type")),
		Centers: asFilterParameter(req.QueryParameter("center"))}

	// TODO optimize
	// if system query parameter is given then first select all systems that match
	// and compute the Centers value of the connection filter.
	systemFilter := req.QueryParameter("system")
	if len(systemFilter) == 0 || !strings.Contains(systemFilter, ":") {
		resp.WriteError(400, errors.New("bad format system query parameter"))
		return
	}
	if len(systemFilter) > 0 {
		all, err := d.service.AllSystems(ctx)
		if err != nil {
			log.Printf("AllSystems failed:%#v", err)
			resp.WriteError(500, err)
			return
		}
		systemAttribute := model.ParseAttribute(systemFilter)
		centers := []string{}
		for _, each := range all {
			if each.HasAttribute(systemAttribute) {
				centers = append(centers, each.ID)
			}
		}
		filter.Centers = centers
	}
	// END optimize

	connections, err := d.service.AllConnections(ctx, filter)
	if err != nil {
		log.Printf("AllConnections failed:%#v", err)
		resp.WriteError(500, err)
		return
	}
	format := req.QueryParameter("format")
	if "" == format {
		format = "svg"
	}
	id, err := model.GenerateUUID()
	if err != nil {
		log.Printf("GenerateUUID failed:%v", err)
		resp.WriteError(500, err)
		return
	}
	input := fmt.Sprintf("%v/%v.dot", DotConfig["tmp"], id)
	output := fmt.Sprintf("%v/%v.%v", DotConfig["tmp"], id, format)

	dotBuilder := application.NewDotBuilder()
	dotBuilder.ClusterBy(req.QueryParameter("cluster"))
	dotBuilder.BuildFromAll(connections)

	dotOnly := req.QueryParameter("format") == "dot"
	if dotOnly {
		resp.AddHeader("Content-Type", "text/plain")
		dotBuilder.WriteDot(resp)
		return
	}
	dotBuilder.WriteDotFile(input)

	cmd := exec.Command(DotConfig["binpath"],
		fmt.Sprintf("-T%v", format),
		fmt.Sprintf("-o%v", output),
		input)
	err = cmd.Start()
	if err != nil {
		log.Printf("Dot command start failed:%v", err)
		resp.WriteError(500, err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Dot did not complete:%v", err)
		resp.WriteError(500, err)
		return
	}
	// resp.AddHeader("Content-Type", "image/svg+xml")
	http.ServeFile(resp, req.Request, output)
}
