package api

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "snaga-team/controllers"
)

func init() {
  r := mux.NewRouter()
  apiRouter := r.PathPrefix("/api").Subrouter()
  apiRouter.HandleFunc("/", handler)
  userRoutes := apiRouter.PathPrefix("/users").Subrouter()
  controllers.InitUserControllerHandlers(userRoutes)

  shipRoutes := apiRouter.PathPrefix("/ships").Subrouter()
  controllers.InitShipControllerHandlers(shipRoutes)

  http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, world from MUX!")
}
