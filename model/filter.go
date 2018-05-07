package model

import "fmt"

// For querying connections ; each field can be single or comma separated of regular expressions
type ConnectionsFilter struct {
	Froms, Tos, Types, Centers []string
}

func (f ConnectionsFilter) String() string {
	return fmt.Sprintf("%#v", f)
}

func (f ConnectionsFilter) Matches(c Connection) bool {
	hasFrom := len(f.Froms) == 0 || contains(f.Froms, c.From)
	hasTo := len(f.Tos) == 0 || contains(f.Tos, c.To)
	hasType := len(f.Types) == 0 || contains(f.Types, c.Type)
	inCenter := len(f.Centers) == 0 || contains(f.Centers, c.From) || contains(f.Centers, c.To)
	return (hasFrom || hasTo || inCenter) && hasType
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
