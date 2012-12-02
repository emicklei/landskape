package dao

import (
	"github.com/emicklei/landskape/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ApplicationDao struct {
	Collection *mgo.Collection
}

func (self ApplicationDao) Save(app *model.Application) error {
	_, err := self.Collection.Upsert(bson.M{"_id": app.Id}, app) // ChangeInfo
	return err
}

func (self ApplicationDao) FindAll() ([]model.Application, error) {
	query := bson.M{}
	result := []model.Application{}
	err := self.Collection.Find(query).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (self ApplicationDao) FindById(id string) (model.Application, error) {
	result := model.Application{}
	err := self.Collection.FindId(id).One(&result)
	return result, err
}
