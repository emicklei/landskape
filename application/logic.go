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
	SystemDao     dao.SystemDataAccess
	ConnectionDao dao.ConnectionDataAccess
}

func (l Logic) AllSystems(ctx context.Context) ([]model.System, error) {
	apps, err := l.SystemDao.FindAll(ctx)
	if err != nil {
		return []model.System{}, err
	}
	return apps, nil
}

func (l Logic) AllConnections(ctx context.Context, filter model.ConnectionsFilter) ([]model.Connection, error) {
	cons, err := l.ConnectionDao.FindAllMatching(ctx, filter)
	if err != nil {
		return []model.Connection{}, err
	}
	sys := map[string]model.System{}
	// populate all systems, sequential for now
	resolved := []model.Connection{}
	for _, each := range cons {
		other := model.Connection{
			Journal:    each.Journal,
			To:         each.To,
			From:       each.From,
			Type:       each.Type,
			Attributes: each.Attributes,
		}
		from, ok := sys[each.From]
		if !ok {
			s, err := l.GetSystem(ctx, each.From)
			if err != nil {
				return []model.Connection{}, err
			}
			from = s
			sys[each.From] = from
		}
		other.FromSystem = from
		to, ok := sys[each.To]
		if !ok {
			s, err := l.GetSystem(ctx, each.To)
			if err != nil {
				return []model.Connection{}, err
			}
			to = s
			sys[each.To] = to
		}
		other.ToSystem = to
		resolved = append(resolved, other)
	}
	return resolved, nil
}

func (l Logic) DeleteConnection(ctx context.Context, con model.Connection) error {
	return l.ConnectionDao.Remove(ctx, con)
}

func (l Logic) SaveConnection(ctx context.Context, con model.Connection, createIfAbsent bool) error {
	if len(con.From) == 0 {
		return errors.New("Invalid from (empty or non-exist):" + con.From)
	}
	if !l.ExistsSystem(ctx, con.From) {
		if createIfAbsent {
			_, err := l.SaveSystem(ctx, model.NewSystem(con.From))
			if err != nil {
				return err
			}
		}
	}
	if len(con.To) == 0 {
		return errors.New("Invalid to (empty or non-exist):" + con.To)
	}
	if !l.ExistsSystem(ctx, con.To) {
		if createIfAbsent {
			_, err := l.SaveSystem(ctx, model.NewSystem(con.To))
			if err != nil {
				return err
			}
		}
	}
	// TODO move to validate
	if len(con.Type) == 0 {
		return errors.New("Invalid type (empty)")
	}
	return l.ConnectionDao.Save(ctx, con)
}

func (l Logic) GetSystem(ctx context.Context, id string) (model.System, error) {
	return l.SystemDao.FindById(ctx, id)
}

func (l Logic) DeleteSystem(ctx context.Context, id string) error {
	// TODO remove all its connections
	return l.SystemDao.RemoveById(ctx, id)
}

func (l Logic) ExistsSystem(ctx context.Context, id string) bool {
	return l.SystemDao.Exists(ctx, id)
}

func (l Logic) SaveSystem(ctx context.Context, app *model.System) (*model.System, error) {
	app.Modified = time.Now()
	return app, l.SystemDao.Save(ctx, app)
}

func (l Logic) ChangeSystemId(ctx context.Context, oldId, newId string) (*model.System, error) {
	target, err := l.GetSystem(ctx, oldId)
	if err != nil {
		return nil, errors.New("No such system:" + oldId)
	}
	_, err = l.GetSystem(ctx, newId)
	if err == nil {
		return nil, errors.New("System already exists:" + newId)
	}
	newSystem := model.NewSystem(newId)
	newSystem.Attributes = target.Attributes
	return l.SaveSystem(ctx, newSystem)
}
