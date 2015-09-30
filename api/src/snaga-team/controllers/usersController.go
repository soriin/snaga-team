package controllers

import (
  "appengine"
  "appengine/datastore"

  "github.com/gorilla/mux"

  "net/http"
  "strings"

  "snaga-team/helpers"
  "snaga-team/models"
)

func InitUserControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allUsers).Methods("GET")
  router.HandleFunc("/", addUser).Methods("POST")
  router.HandleFunc("/{id}", updateUser).Methods("PUT")
  router.HandleFunc("/", deleteUser).Methods("DELETE")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processAllUsers(c, w, r)
}

func processAllUsers(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  _, err := helpers.VerifyGoogleToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  q := datastore.NewQuery("user")
  var users []models.User

  for t := q.Run(c); ; {
    var x models.User
    key, err := t.Next(&x)

    if err == datastore.Done {
      break
    }
    if err != nil {
      helpers.SendError(w, err.Error(), http.StatusInternalServerError)
      return
    }
    x.Id = key.Encode()
    users = append(users, x)
  }

  helpers.SendJson(w, users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processAddUser(c, w, r)
}

func processAddUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  tokenEmail, err := helpers.VerifyGoogleToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  var newUser models.User
  thisUser, err := getUserWithEmail(c, tokenEmail)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if thisUser != nil {
    helpers.SendError(w, "", 409)
    return
  }

  err = helpers.ReadJson(r.Body, &newUser)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  newUser.Email = tokenEmail
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "user", nil), &newUser)
  newUser.Id = key.Encode()

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = helpers.SendJson(w, newUser)

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  }
}

func updateUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processUpdateUser(c, w, r)
}

func processUpdateUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  var newUser models.User
  id := mux.Vars(r)["id"]
  tokenEmail, err := helpers.VerifyGoogleToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = helpers.ReadJson(r.Body, &newUser)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  myKey, err := datastore.DecodeKey(id)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusNotFound)
    return
  }

  var currentUserData models.User
  err = datastore.Get(c, myKey, &currentUserData)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if strings.ToLower(tokenEmail) != strings.ToLower(currentUserData.Email) {
    helpers.SendError(w, "User not authorized to modify this user's data", http.StatusForbidden)
    return
  }

  key, err := datastore.Put(c, myKey, &newUser)
  newUser.Id = key.Encode()

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = helpers.SendJson(w, newUser)

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  }
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processDeleteUser(c, w, r)
}

func processDeleteUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {

}

func getUserWithEmail(c appengine.Context, email string) (*models.User, error) {
  q := datastore.NewQuery("user").Filter("Email =", email)

  var x models.User
  key, err := q.Run(c).Next(&x)

  if err == datastore.Done {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }
  x.Id = key.Encode()
  return &x, nil
}
