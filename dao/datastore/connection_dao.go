package datastore

import (
	"context"

	"github.com/emicklei/landskape/model"
)

type ConnectionDao struct {
}

func (s ConnectionDao) FindAllMatching(ctx context.Context, scope string, filter model.ConnectionsFilter) ([]model.Connection, error) {
	return []model.Connection{}, nil
}
func (s ConnectionDao) Save(ctx context.Context, con model.Connection) error {
	return nil
}
func (s ConnectionDao) Remove(ctx context.Context, con model.Connection) error {
	return nil
}
func (s ConnectionDao) RemoveAllToOrFrom(ctx context.Context, scope, toOrFrom string) error {
	return nil
}
