package application

import (
	"github.com/emicklei/landskape/dao"
	"github.com/emicklei/landskape/model"
)

type Logic struct {
	applicationDao dao.ApplicationDao
	connectionDao  dao.ConnectionDao
}

func (self Logic) AllApplications() (model.Applications, error) {
	apps := []model.Application{}
	apps = append(apps, model.Application{})
	apps = append(apps, model.Application{})
	return model.Applications{apps}, nil
}
