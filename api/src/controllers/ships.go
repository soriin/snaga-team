package controllers

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func InitShipControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allShips)
}

func allShips(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "All Ships!")
}
