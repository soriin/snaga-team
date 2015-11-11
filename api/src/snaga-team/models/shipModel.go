package models

type Ship struct {
	Id string `datastore:"-"`
	DisplayName string
	Manufacturer string
	IconImageUrl string
}
