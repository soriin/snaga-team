package models

type Event struct {
  Id string `datastore:"-"`

  // Identifying information
  Title string
  CreatorId string
  LocationId string
  LocationDetails string
  Ships []EventShip

  // Dates
  CreateDate string //TODO: Figure out datetime type
  EventDate string
  CancellationDate string
  LastUpdateDate string

  // Permissions
  Whitelist []string
  Blacklist []string
}
