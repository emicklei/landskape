package datastore

import (
	"context"

	"github.com/emicklei/landskape/model"
)

type SystemDao struct {
}

func (s SystemDao) Exists(ctx context.Context, scope, id string) bool {
	return false
}
func (s SystemDao) Save(ctx context.Context, app *model.System) error {
	return nil
}
func (s SystemDao) FindAll(ctx context.Context, scope string) ([]model.System, error) {
	return []model.System{}, nil
}
func (s SystemDao) FindById(ctx context.Context, scope, id string) (model.System, error) {
	return model.System{}, nil
}
func (s SystemDao) RemoveById(ctx context.Context, scope, id string) error {
	return nil
}
