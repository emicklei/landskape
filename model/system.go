package model

import (
	"cloud.google.com/go/datastore"
)

// System is the generic name for a IT landscape object.
// Examples are: Webservice, Database schema, Ftp server, Third party solution
type System struct {
	// internal
	DBKey      *datastore.Key `datastore:"__key__" json:"-"`
	Attributes []Attribute    `datastore:",flatten" json:"attributes"`

	// populated from DBKey
	ID string `datastore:"-" json:"id"`
}

func NewSystem(id string) *System {
	return &System{DBKey: NewSystemKey(id)}
}

func NewSystemKey(id string) *datastore.Key {
	key := datastore.NameKey("System", id, nil)
	key.Namespace = "landskape"
	return key
}

// AttributeList exists for AttributesHolder
func (s System) AttributeList() []Attribute { return s.Attributes }

func (s *System) DeleteAttribute(name string) {
	s.SetAttribute(name, "")
}

func (s *System) SetAttribute(name, value string) {
	if len(name) == 0 {
		return
	}
	if len(value) == 0 {
		// remove it
		without := []Attribute{}
		for _, each := range s.Attributes {
			if each.Name != name {
				without = append(without, each)
			}
		}
		s.Attributes = without
		return
	}
	// replace or add
	for i, each := range s.Attributes {
		if each.Name == name {
			s.Attributes[i] = Attribute{Name: name, Value: value}
			return
		}
	}
	// not found, add it
	s.Attributes = append(s.Attributes, Attribute{Name: name, Value: value})
}

// HasAttribute returns whether it has an attribute with the same name and value.
func (s System) HasAttribute(a Attribute) bool {
	for _, each := range s.Attributes {
		if each.Name == a.Name && each.Value == a.Value {
			return true
		}
	}
	return false
}
