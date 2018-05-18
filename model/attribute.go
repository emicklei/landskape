package model

import "strings"

// Attribute is a generic key-value pair of strings
// Each attribute has its own lifecyle to track value changes
type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

// For querying Systems and Connections ; each field can be a regular expression
type AttributesFilter struct {
	Name, Value string
}

type AttributesHolder interface {
	AttributeList() []Attribute
}

func ParseAttribute(nameColonValue string) Attribute {
	colon := strings.Index(nameColonValue, ":")
	return Attribute{Name: nameColonValue[0:colon], Value: nameColonValue[colon+1:]}
}
