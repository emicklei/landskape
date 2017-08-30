package tiedot

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/emicklei/landskape/model"
)

type SystemDao struct {
	Collection *db.Col
}

func (s SystemDao) Exists(scope, id string) bool {
	return false
}
func (s SystemDao) Save(app *model.System) error {
	return nil
}
func (s SystemDao) FindAll(scope string) ([]model.System, error) {
	return []model.System{}, nil
}
func (s SystemDao) FindById(scope, id string) (model.System, error) {
	return model.System{}, nil
}
func (s SystemDao) RemoveById(scope, id string) error {
	return nil
}
