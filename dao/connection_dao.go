package dao

import (
	"context"
	"log"
	"time"

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
	for _, each := range list {
		each.Attributes = append(each.Attributes, model.Attribute{Name: "id", Value: each.DBKey.Name})
	}
	return list, err
}
func (s ConnectionDao) Save(ctx context.Context, con model.Connection) error {
	log.Printf("saving connection:%#v\n", con)
	if con.DBKey == nil {
		id, _ := model.GenerateUUID()
		key := datastore.NameKey(conKind, id, nil)
		con.DBKey = key
	}
	con.Journal.Modified = time.Now()
	_, err := s.client.Put(ctx, con.DBKey, &con)
	return err
}
func (s ConnectionDao) Remove(ctx context.Context, con model.Connection) error {
	return s.client.Delete(ctx, con.DBKey)
}
func (s ConnectionDao) RemoveAllToOrFrom(ctx context.Context, toOrFrom string) error {
	return nil
}
