package application

import (
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/model"
)

type Logic struct {
	ApplicationDao dao.ApplicationDao
	ConnectionDao  dao.ConnectionDao
}

func (self Logic) AllApplications() (model.Applications, error) {
	cons := []model.Application{}
	cons = append(cons, model.Application{})
	cons = append(cons, model.Application{})
	return model.Applications{cons}, nil
}

func (self Logic) AllConnections() (model.Connections, error) {
	cons := []model.Connection{}
	cons = append(cons, model.Connection{})
	cons = append(cons, model.Connection{})
	return model.Connections{cons}, nil
}
