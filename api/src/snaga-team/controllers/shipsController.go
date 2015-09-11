package controllers

import (
  "net/http"
  // "reflect"

  "github.com/gorilla/mux"

  "snaga-team/models"
  "snaga-team/helpers"
  // "snaga-team/services/datarepo"

  "appengine"
  "appengine/datastore"
)

func InitShipControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allShips).Methods("GET")
  router.HandleFunc("/", addShip).Methods("POST")
}

func allShips(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  q := datastore.NewQuery("ship")
  var ships []models.Ship

  for t := q.Run(c); ; {
    var x models.Ship
    key, err := t.Next(&x)

    if err == datastore.Done {
      break
    }
    if err != nil {
      helpers.SendError(w, err.Error(), http.StatusInternalServerError)
      return
    }
    x.Id = key.Encode()
    ships = append(ships, x)
  }

  helpers.SendJson(w, ships)
}

func addShip(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  var newShip models.Ship

  err := helpers.ReadJson(r, &newShip)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "ship", nil), &newShip)
  newShip.Id = key.Encode()

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = helpers.SendJson(w, newShip)

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  }
}
