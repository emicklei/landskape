package model

import (
	"time"
)

type Validator interface {
	Validate() error
}

type AttributesHolder interface {
	AttributeList() []Attribute
}

// Journal is to track who (or what System)
// is responsible for the current state of the containing struct.
type Journal struct {
	Modified   time.Time
	ModifiedBy string
}

// System is the generic name for a IT landscape object.
// Examples are: Webservice, Database schema, Ftp server, Third party solution
type System struct {
	Journal
	ID         string
	Attributes []Attribute
}

func (s System) AttributeList() []Attribute { return s.Attributes }

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
	for _, each := range s.Attributes {
		if each.Name == name {
			each.Value = value
			return
		}
	}
	// not found, add it
	s.Attributes = append(s.Attributes, Attribute{Name: name, Value: value})
}

// Attribute is a generic key-value pair of strings
// Each attribute has its own lifecyle to track value changes
type Attribute struct {
	Journal
	Name, Value string
}

// AttributeValue finds the value of an attribute for a given name, return empty string if not found
func AttributeValue(holder AttributesHolder, name string) string {
	for _, each := range holder.AttributeList() {
		if each.Name == name {
			return each.Value
		}
	}
	return ""
}

// Connection is the generic name for a logical connection between 2 IT landscape object.
// From and To refer to the Id of the System.
// Example of Type are:  http, https, aq, jdbc, ftp, smtp
type Connection struct {
	Journal
	From, To   string
	Type       string
	Attributes []Attribute
	// populated
	FromSystem, ToSystem System
}

func (c Connection) Validate() error {
	return nil // TODO
}

func (c Connection) AttributeList() []Attribute {
	return c.Attributes
}

func (c Connection) ID() string {
	return c.From + "_" + c.Type + "_" + c.To
}

// For querying connections ; each field can be single or comma separated of regular expressions
type ConnectionsFilter struct {
	Froms, Tos, Types, Centers []string
}

// For querying Systems and Connections ; each field can be a regular expression
type AttributesFilter struct {
	Name, Value string
}

// Systems is a container of System for XML/JSON export
type Systems struct{ List []System }

// Connections is a container of System for XML/JSON export
type Connections struct{ List []Connection }
