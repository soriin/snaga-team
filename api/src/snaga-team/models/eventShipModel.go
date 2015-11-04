package models

type EventShip struct {
  Id string `datastore:"-"`
  ShipTypeId string
  DesiredCrewCount int
  Crew []CrewMember
  OwnerId string
  AddedDate string
  LastUpdateDate string
}
