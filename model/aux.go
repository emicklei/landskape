package model

import "time"

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

// For querying connections ; each field can be single or comma separated of regular expressions
type ConnectionsFilter struct {
	Froms, Tos, Types, Centers []string
}

// For querying Systems and Connections ; each field can be a regular expression
type AttributesFilter struct {
	Name, Value string
}
