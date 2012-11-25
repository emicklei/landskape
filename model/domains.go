package model

import (
	"time"
)

type Lifecyle struct {
	Created, Updated     time.Time
	CreatedBy, UpdatedBy string
}

type Application struct {
	Lifecyle
	Id         string
	Attributes []Attribute
}
type Attribute struct {
	Lifecyle
	Name, Value string
}
type Connection struct {
	Lifecyle
	From, To, Type string
	Attributes     []Attribute
}

type Applications struct{ Application []Application }
type Connections struct{ Connection []Connection }
