package dao

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/emicklei/landskape/model"
)

const conKind = "landskape.Connection"

type ConnectionDao struct {
	client *datastore.Client
}

func NewConnectionDao(ds *datastore.Client) *ConnectionDao {
	return &ConnectionDao{client: ds}
}

func (s ConnectionDao) FindAllMatching(ctx context.Context, filter model.ConnectionsFilter) ([]model.Connection, error) {
	var list []model.Connection
	query := datastore.NewQuery(conKind)
	_, err := s.client.GetAll(ctx, query, &list)
	return list, err
}
func (s ConnectionDao) Save(ctx context.Context, con model.Connection) error {
	log.Printf("saving connection:%#v\n", con)
	key := datastore.NameKey(conKind, con.ID(), nil)
	_, err := s.client.Put(ctx, key, &con)
	return err
}
func (s ConnectionDao) Remove(ctx context.Context, con model.Connection) error {
	key := datastore.NameKey(conKind, con.ID(), nil)
	return s.client.Delete(ctx, key)
}
func (s ConnectionDao) RemoveAllToOrFrom(ctx context.Context, toOrFrom string) error {
	return nil
}
