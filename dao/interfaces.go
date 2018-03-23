package dao

import (
	"context"

	"github.com/emicklei/landskape/model"
)

type ConnectionDao interface {
	FindAllMatching(ctx context.Context, scope string, filter model.ConnectionsFilter) ([]model.Connection, error)
	Save(ctx context.Context, con model.Connection) error
	Remove(ctx context.Context, con model.Connection) error
	RemoveAllToOrFrom(ctx context.Context, scope, toOrFrom string) error
}

type SystemDao interface {
	Exists(ctx context.Context, scope, id string) bool
	Save(ctx context.Context, app *model.System) error
	FindAll(ctx context.Context, scope string) ([]model.System, error)
	FindById(ctx context.Context, scope, id string) (model.System, error)
	RemoveById(ctx context.Context, scope, id string) error
}
