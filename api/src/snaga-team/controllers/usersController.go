package controllers

import (
  "appengine"
  "appengine/datastore"

  "github.com/gorilla/mux"

  "golang.org/x/net/context"
  "golang.org/x/oauth2/google"

  "google.golang.org/api/oauth2/v2"

  "net/http"

  "snaga-team/helpers"
  "snaga-team/models"
)

func InitUserControllerHandlers(router *mux.Router) {
  router.HandleFunc("/", allUsers).Methods("GET")
  router.HandleFunc("/", addUser).Methods("POST")
  router.HandleFunc("/", deleteUser).Methods("DELETE")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processAllUsers(c, w, r)
}

func processAllUsers(c appengine.Context, w http.ResponseWriter, r *http.Request) {
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
  var newUser models.User
  content := struct{TokenId string}{}

  err := helpers.ReadJson(r.Body, &content)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  ctx := context.Background()
  client, err := google.DefaultClient(ctx, oauth2.UserinfoProfileScope)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  oauth2Service, err := oauth2.New(client)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  tokenSvc := oauth2Service.Tokeninfo()
  tokenSvc = tokenSvc.AccessToken(content.TokenId)
  tokenInfo, err := tokenSvc.Do()
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if tokenInfo.ExpiresIn == 0 {
    return
  }

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

func deleteUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processDeleteUser(c, w, r)
}

func processDeleteUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {

}
