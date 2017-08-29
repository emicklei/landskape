package dao

import "github.com/emicklei/landskape/model"

type ConnectionDao interface {
	FindAllMatching(scope string, filter model.ConnectionsFilter) ([]model.Connection, error)
	Save(con model.Connection) error
	Remove(con model.Connection) error
	RemoveAllToOrFrom(scope, toOrFrom string) error
}

type SystemDao interface {
	Exists(scope, id string) bool
	Save(app *model.System) error
	FindAll(scope string) ([]model.System, error)
	FindById(scope, id string) (model.System, error)
	RemoveById(scope, id string) error
}
