package application

import (
	"github.com/emicklei/landskape/model"
	"testing"
)

func TestDotBuilderOneConnection(t *testing.T) {
	c := []model.Connection{model.Connection{From: "A", To: "B", Type: "T"}}
	b := newDotBuilder()
	b.BuildFromAll(c)
	b.writeDotFile("/tmp/TestDotBuilderOneConnection.dot")
}
