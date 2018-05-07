package model

import "testing"

func TestFilterMatches1(t *testing.T) {
	c := Connection{From: "from", To: "to"}
	f := ConnectionsFilter{Froms: []string{"from"}, Tos: []string{"to"}}
	if !f.Matches(c) {
		t.Error("should match")
	}
}

func TestFilterMatchesNoFroms(t *testing.T) {
	c := Connection{From: "from", To: "to"}
	f := ConnectionsFilter{Tos: []string{"to"}}
	if !f.Matches(c) {
		t.Error("should match")
	}
}

func TestFilterMatchesNoTos(t *testing.T) {
	c := Connection{From: "from", To: "to"}
	f := ConnectionsFilter{Froms: []string{"from"}}
	if !f.Matches(c) {
		t.Error("should match")
	}
}
func TestFilterMatchesWithType(t *testing.T) {
	c := Connection{From: "from", To: "to", Type: "type"}
	f := ConnectionsFilter{Types: []string{"type", "epyt"}}
	if !f.Matches(c) {
		t.Error("should match")
	}
}
func TestFilterMatchesWithWrongType(t *testing.T) {
	c := Connection{From: "from", To: "to", Type: "type"}
	f := ConnectionsFilter{Types: []string{"epyt"}}
	if f.Matches(c) {
		t.Error("should not match")
	}
}
