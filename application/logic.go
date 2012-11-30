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

func (self Logic) ExistsApplication(id string) bool {
	result, _ := self.ApplicationDao.FindById(id)
	log.Printf("res:%#v", result)
	return result.Id == id
}

func (self Logic) SaveApplication(app *model.Application) (*model.Application, error) {
	return app, self.ApplicationDao.Save(app)
}
