package controllers

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "models"
)

func InitShipControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allShips).Methods("GET")
  router.HandleFunc("/", addShip).Methods("POST")
}

func allShips(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "All Ships!")
}

func addShip(w http.ResponseWriter, r *http.Request) {
  newShip := models.Ship{}
  newShip.DisplayName = "test ship"

  jStream, err := json.Marshal(newShip)
  if err != nil {
    fmt.Fprint(w, err)
  }

  fmt.Fprint(w, string(jStream[:]))
}
