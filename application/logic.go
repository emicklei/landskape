package application

import (
	"context"
	"errors"

	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/model"
	//	"log"
	"time"
)

type Logic struct {
	SystemDao     dao.SystemDao
	ConnectionDao dao.ConnectionDao
}

func (l Logic) AllSystems(ctx context.Context, scope string) (model.Systems, error) {
	apps, err := l.SystemDao.FindAll(ctx, scope)
	if err != nil {
		return model.Systems{}, err
	}
	return model.Systems{apps}, nil
}

func (l Logic) AllConnections(ctx context.Context, scope string, filter model.ConnectionsFilter) (model.Connections, error) {
	cons, err := l.ConnectionDao.FindAllMatching(ctx, scope, filter)
	if err != nil {
		return model.Connections{}, err
	}
	return model.Connections{cons}, nil
}

func (l Logic) DeleteConnection(ctx context.Context, con model.Connection) error {
	return l.ConnectionDao.Remove(ctx, con)
}

func (l Logic) SaveConnection(ctx context.Context, con model.Connection) error {
	// Check from and to for existence
	if con.From == "" || !l.ExistsSystem(ctx, con.Scope, con.From) {
		return errors.New("Invalid from (empty or non-exist):" + con.From)
	}
	if con.To == "" || !l.ExistsSystem(ctx, con.Scope, con.To) {
		return errors.New("Invalid to (empty or non-exist):" + con.To)
	}
	if con.Type == "" {
		return errors.New("Invalid type (empty)")
	}
	return l.ConnectionDao.Save(ctx, con)
}

func (l Logic) GetSystem(ctx context.Context, scope, id string) (model.System, error) {
	return l.SystemDao.FindById(ctx, scope, id)
}

func (l Logic) DeleteSystem(ctx context.Context, scope, id string) error {
	// TODO remove all its connections
	return l.SystemDao.RemoveById(ctx, scope, id)
}

func (l Logic) ExistsSystem(ctx context.Context, scope, id string) bool {
	return l.SystemDao.Exists(ctx, scope, id)
}

func (l Logic) SaveSystem(ctx context.Context, app *model.System) (*model.System, error) {
	app.Modified = time.Now()
	return app, l.SystemDao.Save(ctx, app)
}

func (l Logic) ChangeSystemId(ctx context.Context, scope, oldId, newId string) (*model.System, error) {
	target, err := l.GetSystem(ctx, scope, oldId)
	if err != nil {
		return nil, errors.New("No such system:" + oldId + " in scope:" + scope)
	}
	_, err = l.GetSystem(ctx, scope, newId)
	if err == nil {
		return nil, errors.New("System already exists:" + newId + " in scope:" + scope)
	}
	newSystem := &model.System{ID: newId, Scope: scope}
	newSystem.Attributes = target.Attributes
	return l.SaveSystem(ctx, newSystem)
}
