package dao

import (
	"github.com/emicklei/landskape/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ConnectionDao struct {
	Collection *mgo.Collection
}

func (self ConnectionDao) FindAllMatching(filter model.ConnectionsFilter) ([]model.Connection, error) {
	query := bson.M{}
	result := []model.Connection{}
	err := self.Collection.Find(query).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (self ConnectionDao) Save(con model.Connection) error {
	query := bson.M{"from": con.From, "to": con.To, "type": con.Type}
	_, err := self.Collection.Upsert(query, con) // ChangeInfo
	return err
}

func (self ConnectionDao) Remove(con model.Connection) error {
	query := bson.M{"from": con.From, "to": con.To, "type": con.Type}
	return self.Collection.Remove(query)
}
