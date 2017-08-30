package tiedot

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/emicklei/landskape/model"
)

type SystemDao struct {
	Systems *db.Col
}

func (s SystemDao) Exists(scope, id string) bool {
	return false
}
func (s SystemDao) Save(app *model.System) error {
	doc := map[string]interface{}{
		"id":         app.Id,
		"scope":      app.Scope,
		"modified":   app.Journal.Modified,
		"modifiedBy": app.Journal.ModifiedBy,
	}
	for _, each := range app.AttributeList() {
		doc[each.Name] = each.Value
	}
	if app.DatabaseID == 0 {
		dbid, err := s.Systems.Insert(doc)
		app.DatabaseID = dbid
		return err
	} else {
		return s.Systems.Update(app.DatabaseID, doc)
	}
}
func (s SystemDao) FindAll(scope string) ([]model.System, error) {
	return []model.System{}, nil
}
func (s SystemDao) FindById(scope, id string) (model.System, error) {
	var query interface{}
	sys := model.System{}
	json.Unmarshal([]byte(`[{"eq": scope, "in": ["scope"]}, {"eq": id, "in": ["id"]}]`), &query)
	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys
	if err := db.EvalQuery(query, s.Systems, &queryResult); err != nil {
		return sys, err
	}
	// Query result are document IDs
	for id := range queryResult {
		// To get query result document, simply read it
		readBack, err := s.Systems.Read(id)
		if err != nil {
			return sys, err
		}
		sys.Id = readBack["id"].(string)
		sys.Scope = readBack["scope"].(string)
		sys.Modified = readBack["modified"].(time.Time)
		sys.ModifiedBy = readBack["modifiedBy"].(string)
		return sys, nil
	}
	return sys, errors.New("no such system")
}
func (s SystemDao) RemoveById(scope, id string) error {
	sys, err := s.FindById(scope, id)
	if err != nil {
		return err
	}
	return s.Systems.Delete(sys.DatabaseID)
}
