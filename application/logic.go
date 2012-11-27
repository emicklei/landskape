package application

import (
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/model"
	"log"
)

var SharedLogic Logic

type Logic struct {
	ApplicationDao dao.ApplicationDao
	ConnectionDao  dao.ConnectionDao
}

func (self Logic) AllApplications() (model.Applications, error) {
	log.Printf("[landskape] get them all")
	apps, err := self.ApplicationDao.FindAll()
	if err != nil {
		return model.Applications{}, err
	}
	return model.Applications{apps}, nil
}

func (self Logic) AllConnections() (model.Connections, error) {
	cons := []model.Connection{}
	cons = append(cons, model.Connection{})
	cons = append(cons, model.Connection{})
	return model.Connections{cons}, nil
}

func (self Logic) SaveApplication(app *model.Application) (*model.Application, error) {
	log.Printf("[landskape] saving %#v", app)
	return app, self.ApplicationDao.Save(app)
}
