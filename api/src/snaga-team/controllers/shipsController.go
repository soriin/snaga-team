package controllers

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"

  "snaga-team/models"
  "snaga-team/services/datarepo"

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

  repo := datarepo.NewDataRepo(c)
  key, err := repo.Put(&newShip, "ship", nil)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  var savedShip models.Ship
  if err = datastore.Get(c, key, &savedShip); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  jStream, err := json.Marshal(newShip)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  fmt.Fprint(w, string(jStream[:]))
}
