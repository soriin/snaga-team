package models

type CrewMember struct {
  Id string `datastore:"-"`
  UserId string
  UserDisplayName string
  DesiredPosition string
  AssignedPosition string
  AddedDate string
  LastUpdateDate string
  IsAccepted bool
}
