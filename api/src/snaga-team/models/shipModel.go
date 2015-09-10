package models

type Ship struct {
	Id string `datastore:"-"`
	DisplayName string
	InGameName string
}
