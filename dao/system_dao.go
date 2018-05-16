package dao

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/emicklei/landskape/model"
)

var systemKind = "System"

type SystemDao struct {
	client *datastore.Client
}

func NewSystemDao(ds *datastore.Client) *SystemDao {
	return &SystemDao{client: ds}
}

func (s SystemDao) Exists(ctx context.Context, id string) bool {
	key := datastore.NameKey(systemKind, id, nil)
	key.Namespace = "landskape"
	return s.client.Get(ctx, key, new(model.System)) == nil
}

func (s SystemDao) Save(ctx context.Context, app *model.System) error {
	_, err := s.client.Put(ctx, app.DBKey, app)
	return err
}

func (s SystemDao) FindAll(ctx context.Context) ([]model.System, error) {
	var list []model.System
	query := datastore.NewQuery(systemKind).Namespace("landskape")
	withTimeout, _ := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	_, err := s.client.GetAll(withTimeout, query, &list)
	return postLoadSystems(list...), err
}

func (s SystemDao) FindById(ctx context.Context, id string) (sys model.System, err error) {
	key := datastore.NameKey(systemKind, id, nil)
	key.Namespace = "landskape"
	err = s.client.Get(ctx, key, &sys)
	if err != nil {
		return sys, err
	}
	sys = postLoadSystems(sys)[0]
	return
}

func (s SystemDao) RemoveById(ctx context.Context, id string) error {
	key := datastore.NameKey(systemKind, id, nil)
	key.Namespace = "landskape"
	return s.client.Delete(ctx, key)
}

func postLoadSystems(systems ...model.System) (list []model.System) {
	for _, each := range systems {
		if each.DBKey != nil {
			each.ID = each.DBKey.Name
			list = append(list, each)
		} else {
			log.Println("ERROR:", each)
		}
	}
	return list
}
