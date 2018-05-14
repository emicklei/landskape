package dao

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"github.com/emicklei/landskape/model"
	"github.com/emicklei/tre"
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
	query := datastore.NewQuery(conKind).Namespace("landskape")
	// for now get all and post filter here
	_, err := s.client.GetAll(ctx, query, &list)
	if err != nil {
		return list, tre.New(err, "GetAll failed", "kind", conKind, "filter", filter)
	}
	filtered := []model.Connection{}
	for _, each := range list {
		if filter.Matches(each) {
			filtered = append(filtered, each)
		}
	}
	return filtered, nil
}

func (s ConnectionDao) Save(ctx context.Context, con model.Connection) error {
	if con.DBKey == nil {
		id, _ := model.GenerateUUID()
		key := datastore.NameKey(conKind, id, nil)
		key.Namespace = "landskape"
		con.DBKey = key
	}
	_, err := s.client.Put(ctx, con.DBKey, &con)
	return err
}

func (s ConnectionDao) Remove(ctx context.Context, con model.Connection) error {
	if con.DBKey == nil {
		return errors.New("nil connection DBKey")
	}
	return s.client.Delete(ctx, con.DBKey)
}

func (s ConnectionDao) RemoveAllToOrFrom(ctx context.Context, toOrFrom string) error {
	return nil
}
