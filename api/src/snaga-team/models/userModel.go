package models

type User struct {
	Id string `datastore:"-"`
	DisplayName string
	InGameName string
	FirstName string
	LastName string
	Email string
	Groups []string
}
