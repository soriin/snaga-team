package models

import (
	"appengine/datastore"
)

type User struct {
	Id string `datastore:"-"`
	DisplayName string
	InGameName string
	FirstName string
	LastName string
	Email string
	Groups []string
	IsAdmin bool `datastore:"-"`
}

func (x *User) Load(c <-chan datastore.Property) error {
    // Load saved properties as usual
    if err := datastore.LoadStruct(x, c); err != nil {
        return err
    }
    // Derive the IsAdmin field.
    x.IsAdmin = x.IsSystemAdmin()
    return nil
}

func (x *User) Save(c chan<- datastore.Property) error {
	return datastore.SaveStruct(x, c)
}
