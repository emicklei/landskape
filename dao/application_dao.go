package dao

import (
	"github.com/emicklei/landskape/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

//	"log"
)

type ApplicationDao struct {
	Collection *mgo.Collection
}

func (self ApplicationDao) Save(app *model.Application) error {
	return self.Collection.Insert(app)
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
	query := bson.M{"Id": id}
	result := model.Application{}
	err := self.Collection.Find(query).One(&result)
	return result, err
}
