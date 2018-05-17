package model

import "testing"

func TestParseAttribute(t *testing.T) {
	a := ParseAttribute("key:value")
	if a.Name != "key" {
		t.Fail()
	}
	if a.Value != "value" {
		t.Fail()
	}
}
