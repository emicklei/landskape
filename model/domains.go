package model

import (
	"time"
)

type Lifecyle struct {
	Created, Updated time.Time
	Who	string
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
