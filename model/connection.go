package model

import "cloud.google.com/go/datastore"

// Connection is the generic name for a logical connection between 2 IT landscape object.
// From and To refer to the Id of the System.
// Example of Type are:  http, https, aq, jdbc, ftp, smtp
type Connection struct {
	// internal
	DBKey *datastore.Key `datastore:"__key__" json:"-"`

	From, To   string
	Type       string      `datastore:"Type,noindex"`
	Attributes []Attribute `datastore:",flatten"`
	// populated
	FromSystem, ToSystem System `datastore:"-" json:"-"`
}

func (c Connection) Validate() error {
	return nil // TODO
}

func (c Connection) AttributeList() []Attribute {
	return c.Attributes
}

func (c *Connection) SetAttribute(name, value string) {
	if len(name) == 0 {
		return
	}
	if len(value) == 0 {
		// remove it
		without := []Attribute{}
		for _, each := range c.Attributes {
			if each.Name != name {
				without = append(without, each)
			}
		}
		c.Attributes = without
		return
	}
	// replace or add
	for i, each := range c.Attributes {
		if each.Name == name {
			c.Attributes[i] = Attribute{Name: name, Value: value}
			return
		}
	}
	// not found, add it
	c.Attributes = append(c.Attributes, Attribute{Name: name, Value: value})
}

func (c *Connection) DeleteAttribute(name string) {
	c.SetAttribute(name, "")
}
