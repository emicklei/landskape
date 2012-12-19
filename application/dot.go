package application

import (
	"fmt"
	"github.com/emicklei/landskape/model"
	"io"
	"os"
	"strings"
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
}

func (self edge) String() string {
	return fmt.Sprintf("%v -> (%v,%v) -> %v", self.from, self.label, self.color, self.to)
}

func (self dotBuilder) BuildFromAll(connections []model.Connection) {
	for _, each := range connections {
		edge := edge{}
		edge.from = labelForNodeIn(each.From, self.nodes)
		edge.to = labelForNodeIn(each.To, self.nodes)
		labelOrEmpty := model.AttributeValue(each, UI_LABEL)
		if "" != labelOrEmpty {
			// detect reference to other attribute
			if strings.HasPrefix("@", labelOrEmpty) {
				referencedLabel := model.AttributeValue(each, labelOrEmpty[1:])
				if "" != referencedLabel {
					edge.label = referencedLabel
				} else {
					edge.label = labelOrEmpty
				}
			} else {
				edge.label = labelOrEmpty
			}
		} else {
			edge.label = each.Type
		}
		edge.color = colorForLabel(edge.label, model.AttributeValue(each, UI_COLOR))
		self.edges = append(self.edges, edge)
	}
}

func (self dotBuilder) writeDotFile(output string) error {
	fo, err := os.Create(output)
	if err != nil {
		return err
	}
	defer fo.Close()
	io.WriteString(fo, "digraph {")
	// nodes
	usedNodeValues := map[string]bool{}
	for _, each := range self.edges {
		usedNodeValues[each.to] = true
		usedNodeValues[each.from] = true
	}
	for each, ok := range usedNodeValues {
		if ok {
			nodeName := keyByValue(self.nodes, each)
			io.WriteString(fo, fmt.Sprintf(" %v [label=\"%v\"]\n", each, nodeName))
		}
	}
	// edges
	for _, each := range self.edges {
		io.WriteString(fo, fmt.Sprintf("\t %v-> %v [label=\"%v\" "))
		if "" != each.color {
			io.WriteString(fo, fmt.Sprintf(",color=\"%v\" "))
		}
		io.WriteString(fo, fmt.Sprintf("]\n"))
	}
	io.WriteString(fo, "}")
	return nil
}

func keyByValue(mss map[string]string, s string) string {
	for key, value := range mss {
		if value == s {
			return key
		}
	}
	return ""
}

func colorForLabel(label, uicolor string) string {
	if uicolor != "" {
		return fmt.Sprintf("#%v", uicolor)
	}
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
	return ""
}

func labelForNodeIn(id string, nodes map[string]string) string {
	label := nodes[id]
	if label == "" {
		label = fmt.Sprintf("n%v", len(nodes))
		nodes[id] = label
	}
	return label
}
