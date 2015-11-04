package models

type Location struct {
  Id string `datastore:"-"`

  Name string
  StarmapUrl string
  ImageUrl string
}
