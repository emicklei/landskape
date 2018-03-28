package dao

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/emicklei/landskape/model"
)

var systemKind = "landskape.System"

type SystemDao struct {
	client *datastore.Client
}

func NewSystemDao(ds *datastore.Client) *SystemDao {
	return &SystemDao{client: ds}
}

func (s SystemDao) Exists(ctx context.Context, id string) bool {
	key := datastore.NameKey(systemKind, id, nil)
	return s.client.Get(ctx, key, new(model.System)) == nil
}
func (s SystemDao) Save(ctx context.Context, app *model.System) error {
	log.Printf("saving system:%#v\n", app)
	key := datastore.NameKey(systemKind, app.ID, nil)
	_, err := s.client.Put(ctx, key, app)
	return err
}
func (s SystemDao) FindAll(ctx context.Context) ([]model.System, error) {
	var list []model.System
	query := datastore.NewQuery(systemKind)
	_, err := s.client.GetAll(ctx, query, &list)
	return list, err
}
func (s SystemDao) FindById(ctx context.Context, id string) (model.System, error) {
	return model.System{}, nil
}
func (s SystemDao) RemoveById(ctx context.Context, id string) error {
	key := datastore.NameKey(systemKind, id, nil)
	return s.client.Delete(ctx, key)
}
