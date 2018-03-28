package dao

import (
	"context"

	"github.com/emicklei/landskape/model"
)

type ConnectionDataAccess interface {
	FindAllMatching(ctx context.Context, filter model.ConnectionsFilter) ([]model.Connection, error)
	Save(ctx context.Context, con model.Connection) error
	Remove(ctx context.Context, con model.Connection) error
	RemoveAllToOrFrom(ctx context.Context, toOrFrom string) error
}

type SystemDataAccess interface {
	Exists(ctx context.Context, id string) bool
	Save(ctx context.Context, app *model.System) error
	FindAll(ctx context.Context) ([]model.System, error)
	FindById(ctx context.Context, id string) (model.System, error)
	RemoveById(ctx context.Context, id string) error
}
