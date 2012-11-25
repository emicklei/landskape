package dao

import (
	"labix.org/v2/mgo"

//	"labix.org/v2/mgo/bson"
)

type ConnectionDao struct {
	Collection *mgo.Collection
}
