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
	Modified   time.Time `json:"-"`
	ModifiedBy string    `json:"-"`
}

// For querying Systems and Connections ; each field can be a regular expression
type AttributesFilter struct {
	Name, Value string
}
