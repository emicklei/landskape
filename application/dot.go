package application

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/emicklei/dot"

	"github.com/emicklei/landskape/model"
)

const (
	UI_LABEL = "ui-label"
	UI_COLOR = "ui-color"
)

type edge struct {
	from, to, label, color string
}

type dotBuilder struct {
	edges                []edge
	nodes                map[string]string
	graph                *dot.Graph
	clusterAttributeName string
	config               map[string]string
}

func NewDotBuilder() dotBuilder {
	builder := dotBuilder{}
	builder.nodes = map[string]string{}
	builder.graph = dot.NewGraph(dot.Directed)
	return builder
}

func (e *dotBuilder) ClusterBy(clusterAttribute string) *dotBuilder {
	e.clusterAttributeName = clusterAttribute
	return e
}

func (e *dotBuilder) Config(c map[string]string) {
	e.config = c
}

func (e edge) String() string {
	return fmt.Sprintf("%v -> (%v,%v) -> %v", e.from, e.label, e.color, e.to)
}

// BuildFromAll composes the edges and nodes from a collection of Connection
func (e *dotBuilder) BuildFromAll(connections []model.Connection) {
	hasAttributesSet := map[string]bool{}
	// first apply attributes
	for _, each := range connections {
		if len(each.FromSystem.ID) == 0 {
			log.Println("ERROR: system without ID")
			return
		}
		fromGraph := e.graphForSystem(each.FromSystem)
		toGraph := e.graphForSystem(each.ToSystem)
		from := fromGraph.Node(each.FromSystem.ID)
		to := toGraph.Node(each.ToSystem.ID)
		if hasSet, ok := hasAttributesSet[each.FromSystem.ID]; !hasSet || !ok {
			hasAttributesSet[each.FromSystem.ID] = true
			e.setUIAttributesForSystem(from.AttributesMap, each.FromSystem)
		}
		if hasSet, ok := hasAttributesSet[each.ToSystem.ID]; !hasSet || !ok {
			hasAttributesSet[each.ToSystem.ID] = true
			e.setUIAttributesForSystem(to.AttributesMap, each.ToSystem)
		}
	}
	for _, each := range connections {
		fromGraph := e.graphForSystem(each.FromSystem)
		toGraph := e.graphForSystem(each.ToSystem)
		from := fromGraph.Node(each.FromSystem.ID)
		to := toGraph.Node(each.ToSystem.ID)
		// use fromGraph for adding the edge; the Edge func will find the common ancestor.
		edge := fromGraph.Edge(from, to)
		e.setUIAttributesForConnection(edge.AttributesMap, each)
	}
}

func (e *dotBuilder) graphForSystem(sys model.System) *dot.Graph {
	if len(e.clusterAttributeName) == 0 {
		// root
		return e.graph
	}
	clusterValue := model.AttributeValue(sys, e.clusterAttributeName)
	if len(clusterValue) == 0 {
		// root
		return e.graph
	}
	return e.graph.Subgraph(clusterValue, dot.ClusterOption{})
}

func (e *dotBuilder) setUIAttributesForSystem(a dot.AttributesMap, s model.System) {
	a.Attr("label", s.ID)                                 // can be overwritten if ui-label was set
	a.Attr("shape", e.config["default.ui-shape"])         // can be overwritten if ui-shape was set
	a.Attr("color", e.config["default.ui-color"])         // can be overwritten if ui-color was set
	a.Attr("style", e.config["default.ui-style"])         // can be overwritten if ui-style was set
	a.Attr("fillcolor", e.config["default.ui-fillcolor"]) // can be overwritten if ui-fillcolor was set
	for _, each := range s.Attributes {
		if strings.HasPrefix(each.Name, "ui-") {
			key := each.Name[3:]
			if len(each.Value) > 0 {
				a.Attr(key, each.Value)
			}
		}
	}
}

func (e *dotBuilder) setUIAttributesForConnection(a dot.AttributesMap, c model.Connection) {
	a.Attr("label", c.Type) // can be overwritten if ui-label was set
	for _, each := range c.Attributes {
		if strings.HasPrefix(each.Name, "ui-") {
			key := each.Name[3:]
			if len(each.Value) > 0 {
				a.Attr(key, each.Value)
			}
		}
	}
}

func (e dotBuilder) WriteDotFile(output string) error {
	//log.Println(e.graph.String())
	fo, err := os.Create(output)
	if err != nil {
		return err
	}
	defer fo.Close()
	e.graph.Write(fo)
	return nil
}

func (d dotBuilder) WriteDot(w io.Writer) {
	d.graph.Write(w)
}
