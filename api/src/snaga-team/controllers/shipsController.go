package controllers

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"

  "snaga-team/models"
  "snaga-team/helpers"

  "appengine"
  "appengine/datastore"
)

func InitShipControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allShips).Methods("GET")
  router.HandleFunc("/", addShip).Methods("POST")
}

func allShips(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "All Ships!")
}

func addShip(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  newShip := models.Ship{}
  newShip.DisplayName = "test ship"

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
