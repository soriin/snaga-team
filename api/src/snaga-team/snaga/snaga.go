package snaga

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "snaga-team/controllers"
)

func init() {
  r := mux.NewRouter()
  r.HandleFunc("/", handler)

  userRoutes := r.PathPrefix("/users").Subrouter()
  controllers.InitUserControllerHandlers(userRoutes)

  shipRoutes := r.PathPrefix("/ships").Subrouter()
  controllers.InitShipControllerHandlers(shipRoutes)

  http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, world from MUX!")
}
