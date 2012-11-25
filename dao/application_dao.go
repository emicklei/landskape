package dao

import (
	"labix.org/v2/mgo"

//	"labix.org/v2/mgo/bson"
)

type ApplicationDao struct {
	Collection *mgo.Collection
}
