package dao

import (
	"github.com/emicklei/landskape/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type SystemDao struct {
	Collection *mgo.Collection
}

func (self SystemDao) Save(app *model.System) error {
	_, err := self.Collection.Upsert(bson.M{"_id": app.Id}, app) // ChangeInfo
	return err
}

func (self SystemDao) FindAll() ([]model.System, error) {
	model.Debug("dao", self)
	query := bson.M{}
	result := []model.System{}
	err := self.Collection.Find(query).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (self SystemDao) FindById(id string) (model.System, error) {
	result := model.System{}
	err := self.Collection.FindId(id).One(&result)
	return result, err
}

func (self SystemDao) RemoveById(id string) error {
	return self.Collection.Remove(bson.M{"_id": id})
}
