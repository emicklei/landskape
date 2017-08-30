package tiedot

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/emicklei/landskape/model"
)

type ConnectionDao struct {
	Connections *db.Col
	Systems     *db.Col
}

func (s ConnectionDao) FindAllMatching(scope string, filter model.ConnectionsFilter) ([]model.Connection, error) {
	return []model.Connection{}, nil
}
func (s ConnectionDao) Save(con model.Connection) error {
	return nil
}
func (s ConnectionDao) Remove(con model.Connection) error {
	return nil
}
func (s ConnectionDao) RemoveAllToOrFrom(scope, toOrFrom string) error {
	return nil
}
