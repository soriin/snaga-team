package controllers

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func InitUserControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allUsers)
}

func allUsers(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "All Users!")
}
