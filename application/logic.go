package application

import (
	"errors"
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/model"
	"log"
	"time"
)

var SharedLogic Logic

type Logic struct {
	ApplicationDao dao.ApplicationDao
	ConnectionDao  dao.ConnectionDao
}

func (self Logic) AllApplications() (model.Applications, error) {
	apps, err := self.ApplicationDao.FindAll()
	if err != nil {
		return model.Applications{}, err
	}
	return model.Applications{apps}, nil
}

func (self Logic) AllConnections(filter model.ConnectionsFilter) (model.Connections, error) {
	cons, err := self.ConnectionDao.FindAllMatching(filter)
	if err != nil {
		return model.Connections{}, err
	}
	return model.Connections{cons}, nil
}

func (self Logic) SaveConnection(con model.Connection) error {
	log.Printf("logic.save:%#v", con)
	// Check from and to for existence
	if con.From == "" || !self.ExistsApplication(con.From) {
		return errors.New("Invalid from (empty or non-exist):" + con.From)
	}
	if con.To == "" || !self.ExistsApplication(con.To) {
		return errors.New("Invalid to (empty or non-exist):" + con.To)
	}
	if con.Type == "" {
		return errors.New("Invalid type (empty)")
	}
	return self.ConnectionDao.Save(con)
}

func (self Logic) GetApplication(id string) (model.Application, error) {
	return self.ApplicationDao.FindById(id)
}

func (self Logic) DeleteApplication(id string) error {
	// TODO remove all its connections
	return self.ApplicationDao.RemoveById(id)
}

func (self Logic) ExistsApplication(id string) bool {
	return false
	//	result, _ := self.ApplicationDao.Exists(id)
	//	return result.Id == id
}

func (self Logic) SaveApplication(app *model.Application) (*model.Application, error) {
	app.Modified = time.Now()
	return app, self.ApplicationDao.Save(app)
}
