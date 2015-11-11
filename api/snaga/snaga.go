package api

import (
  "net/http"
  "github.com/gorilla/mux"
  "snaga-team/controllers"
)

func init() {
  r := mux.NewRouter()
  apiRouter := r.PathPrefix("/api").Subrouter()

  userRoutes := apiRouter.PathPrefix("/users").Subrouter()
  controllers.InitUserControllerHandlers(userRoutes)

  shipRoutes := apiRouter.PathPrefix("/ships").Subrouter()
  controllers.InitShipControllerHandlers(shipRoutes)

  eventRoutes := apiRouter.PathPrefix("/events").Subrouter()
  controllers.InitEventsControllerHandlers(eventRoutes)

  http.Handle("/", r)
}
