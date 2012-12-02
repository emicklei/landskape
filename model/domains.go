package model

import (
	"time"
)

// Lifecyle is to track who (or what application) 
// is responsible for the current state of the containing struct.
type Lifecyle struct {
	Modified   time.Time
	ModifiedBy string
}

// Application is the generic name for a IT landscape object.
// Examples are: Webservice, Database schema, Ftp server, Third party solution
type Application struct {
	Lifecyle
	Scope      string
	Id         string `bson:"_id"`
	Attributes []Attribute
}

// Attribute is a generic key-value pair of strings
// Each attribute has its own lifecyle to track value changes
type Attribute struct {
	Lifecyle
	Name, Value string
}

// Connection is the generic name for a logical connection between 2 IT landscape object.
// From and To refer to the Id of the Application.
// Example of Type are:  http, https, aq, jdbc, ftp, smtp
type Connection struct {
	Lifecyle
	Scope          string
	From, To, Type string
	Attributes     []Attribute
}

// Applications is a container of Application for XML/JSON export
type Applications struct{ Application []Application }

// Connections is a container of Application for XML/JSON export
type Connections struct{ Connection []Connection }
