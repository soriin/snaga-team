package controllers

import (
  "appengine"
  "appengine/datastore"

  "github.com/gorilla/mux"

  "net/http"
  "strings"

  "snaga-team/config"
  "snaga-team/helpers"
  "snaga-team/models"
)

func InitUserControllerHandlers(router *mux.Router) {
  // Get all events
  router.HandleFunc("/", getAllEvents).Method("GET")
  router.HandleFunc("/{id}", getEvent).Method("GET")
  router.HandleFunc("/", createEvent).Method("POST")
  router.HandleFunc("/{id}", updateEvent).Method("PUT")
  router.HandleFunc("/{id}", deleteEvent).Method("DELETE")
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processGetAllEvents(c, w, r)
}

func processGetAllEvents(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  // _, err := helpers.VerifyGoogleToken(c, r)
  // if err != nil {
  //   helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }

  q := datastore.NewQuery("event")
  var events []models.Event

  for t := q.Run(c); ; {
    var x models.Event
    key, err := t.Next(&x)

    if err == datastore.Done {
      break
    }
    if err != nil {
      helpers.SendError(w, err.Error(), http.StatusInternalServerError)
      return
    }
    x.Id = key.Encode()
    users = append(events, x)
  }

  helpers.SendJson(w, events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processGetEvent(c, w, r, tokenVerifier)
}

func processGetEvent(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {
  tokenEmail, err := verifier.VerifyToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusForbidden)
    return
  }

  thisUser, err := getUserWithEmail(c, tokenEmail)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  id := mux.Vars(r)["id"]
  myKey, err := datastore.DecodeKey(id)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusNotFound)
    return
  }

  var thisEvent models.Event
  err = datastore.Get(c, myKey, &thisEvent)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = helpers.SendJson(w, thisEvent)

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  }
}

func createEvent(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processCreateEvent(c, w, r, tokenVerifier)
}

func processCreateEvent(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {

}

func updateEvent(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processUpdateEvent(c, w, r, tokenVerifier)
}

func processUpdateEvent(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {

}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processDeleteEvent(c, w, r, tokenVerifier)
}

func processDeleteEvent(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {

}
