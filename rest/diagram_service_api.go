package rest

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/landskape/application"
)

func NewDiagramService(s application.Logic) *restful.WebService {
	ws := new(restful.WebService)
	d := DiagramResource{service: s}
	tags := []string{"diagrams"}

	ws.Path("/v1/diagrams").
		Produces("text/plain", "application/svg", "image/png")
	ws.Route(ws.GET("/").To(d.computeDiagram).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc(`Compute a graphical diagram with all (filtered) connections for all systems and the given scope`).
		Param(ws.QueryParameter("from", "comma separated list of system ids")).
		Param(ws.QueryParameter("to", "comma separated list of system ids")).
		Param(ws.QueryParameter("type", "comma separated list of known connection types")).
		Param(ws.QueryParameter("center", "comma separated list of system ids")).
		Param(ws.QueryParameter("cluster", "show clusters based on the values of the give system attribute")).
		Param(ws.QueryParameter("system", "format is name:value. Filter systems based on this attribute pair.")).
		Param(ws.QueryParameter("format", "svg (default), png, json, pdf, gif, jpg")).
		Param(ws.QueryParameter("layout", `<ul>
		<li>dot (default) - filter for drawing directed graphs</li>
		<li>neato - filter for drawing undirected graphs</li>
		<li>twopi - filter for radial layouts of graphs</li>
		<li>circo - filter for circular layout of graphs</li>
		<li>fdp - filter for drawing undirected graphs</li>
		<li>sfdp - filter for drawing large undirected graphs</li>
		<li>patchwork - filter for squarified tree maps</li>
		<li>osage - filter for array-based layouts</li>
		</ul>`)))
	return ws
}
