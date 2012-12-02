package model

import (
	"time"
)

type Validator interface {
	Validate() error
}

// Journal is to track who (or what application) 
// is responsible for the current state of the containing struct.
type Journal struct {
	Modified   time.Time
	ModifiedBy string
}

// Application is the generic name for a IT landscape object.
// Examples are: Webservice, Database schema, Ftp server, Third party solution
type Application struct {
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
// From and To refer to the Id of the Application.
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

// For querying connections ; each field can be a regular expression
type ConnectionsFilter struct {
	From, To, Type, Center string
}

// For querying applications and connections ; each field can be a regular expression
type AttributesFilter struct {
	Name, Value string
}

// Applications is a container of Application for XML/JSON export
type Applications struct{ Application []Application }

// Connections is a container of Application for XML/JSON export
type Connections struct{ Connection []Connection }
