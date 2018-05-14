package model

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
