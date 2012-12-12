package model

import (
	"time"
)

type Validator interface {
	Validate() error
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
	Scope      string
	Id         string `bson:"_id"`
	Attributes []Attribute
}

// Attribute is a generic key-value pair of strings
// Each attribute has its own lifecyle to track value changes
type Attribute struct {
	Journal
	Name, Value string
}

// Connection is the generic name for a logical connection between 2 IT landscape object.
// From and To refer to the Id of the System.
// Example of Type are:  http, https, aq, jdbc, ftp, smtp
type Connection struct {
	Journal
	Scope          string
	From, To, Type string
	Attributes     []Attribute
}

func (self Connection) Validate() error {
	return nil // TODO	
}

// For querying connections ; each field can be single or comma separated of regular expressions
type ConnectionsFilter struct {
	Froms, Tos, Types, Centers []string
}

// For querying Systems and connections ; each field can be a regular expression
type AttributesFilter struct {
	Name, Value string
}

// Systems is a container of System for XML/JSON export
type Systems struct{ System []System }

// Connections is a container of System for XML/JSON export
type Connections struct{ Connection []Connection }
