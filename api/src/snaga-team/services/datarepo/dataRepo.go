package datarepo

import (
	"appengine"
	"appengine/datastore"
	"snaga-team/models"
)

type DataRepo struct {
	MyContext appengine.Context
}

func NewDataRepo(c appengine.Context) *DataRepo {
	return &DataRepo{MyContext: c}
}

func (repo *DataRepo) Put(obj interface{}, kind string, ancestorPath *datastore.Key) (*datastore.Key, error) {
		key, err := datastore.Put(repo.MyContext, datastore.NewIncompleteKey(repo.MyContext, kind, ancestorPath), obj)

		switch x := obj.(type) {
		case *models.Ship:
			x.Id = key.Encode()
		}

		return key, err
}
