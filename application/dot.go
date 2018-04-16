package application

import (
	"fmt"
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
	edges []edge
	nodes map[string]string
	graph *dot.Graph
}

func NewDotBuilder() dotBuilder {
	builder := dotBuilder{}
	builder.nodes = map[string]string{}
	builder.graph = dot.NewDigraph()
	return builder
}

func (e edge) String() string {
	return fmt.Sprintf("%v -> (%v,%v) -> %v", e.from, e.label, e.color, e.to)
}

// BuildFromAll composes the edges and nodes from a collection of Connection
func (e *dotBuilder) BuildFromAll(connections []model.Connection) {
	hasAttributesSet := map[string]bool{}
	for _, each := range connections {
		if len(each.FromSystem.ID()) == 0 {
			log.Printf("%#v", each)
			panic("jammer")
		}
		from := e.graph.Node(each.FromSystem.ID())
		if hasSet, ok := hasAttributesSet[each.FromSystem.ID()]; !hasSet || !ok {
			hasAttributesSet[each.FromSystem.ID()] = true
			setUIAttributesForSystem(from.AttributesMap, each.FromSystem)
		}
		to := e.graph.Node(each.ToSystem.ID())
		if hasSet, ok := hasAttributesSet[each.ToSystem.ID()]; !hasSet || !ok {
			hasAttributesSet[each.ToSystem.ID()] = true
			setUIAttributesForSystem(to.AttributesMap, each.ToSystem)
		}
		edge := e.graph.Edge(from, to)
		setUIAttributesForConnection(edge.AttributesMap, each)
	}
}

func setUIAttributesForSystem(a dot.AttributesMap, s model.System) {
	a.Attr("label", s.ID) // can be overwritten is ui-label was set
	a.Attr("color", colorForLabel(s.ID()))
	for _, each := range s.Attributes {
		if strings.HasPrefix(each.Name, "ui-") {
			key := each.Name[3:]
			if len(each.Value) > 0 {
				a.Attr(key, each.Value)
			}
		}
	}
}

func setUIAttributesForConnection(a dot.AttributesMap, c model.Connection) {
	a.Attr("label", c.Type) // can be overwritten is ui-label was set
	a.Attr("color", colorForLabel(c.Type))
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
	log.Println(e.graph.String())
	fo, err := os.Create(output)
	if err != nil {
		return err
	}
	defer fo.Close()
	e.graph.Write(fo)
	return nil
}

func colorForLabel(label string) string {
	switch {
	case label == "jdbc":
		return "#E01BD0"
	case label == "dblink":
		return "#FF0000"
	case label == "aq":
		return "#0000FF"
	case label == "http":
		return "#EDB845"

	}
	return "#222"
}
