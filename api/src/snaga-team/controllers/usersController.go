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
  router.HandleFunc("/{id}", getUser).Methods("GET")
  router.HandleFunc("/", addUser).Methods("POST")
  router.HandleFunc("/{id}", updateUser).Methods("PUT")
  router.HandleFunc("/", deleteUser).Methods("DELETE")

  router.HandleFunc("/{id}/groups", updateUserGroups).Methods("PUT")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  processAllUsers(c, w, r)
}

func processAllUsers(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  // _, err := helpers.VerifyGoogleToken(c, r)
  // if err != nil {
  //   helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }

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

func getUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processGetUser(c, w, r, tokenVerifier)
}

func processGetUser(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {
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

  if thisUser == nil {
    // Requester doesn't have an account/isn't logged in.
    helpers.SendError(w, "Must be logged in to view users.", http.StatusForbidden)
    return
  }

  // At this point, the requester is a valid user.

  id := mux.Vars(r)["id"]
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
  err = helpers.SendJson(w, currentUserData)

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  }
}

func addUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processAddUser(c, w, r, tokenVerifier)
}

func processAddUser(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {
  tokenEmail, err := verifier.VerifyToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusForbidden)
    return
  }

  var newUser models.User
  thisUser, err := getUserWithEmail(c, tokenEmail)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if thisUser != nil {
    err = helpers.SendJson(w, thisUser)

    if err != nil {
      helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    }
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
  tokenVerifier := helpers.GetTokenVerifier(r)
  processUpdateUser(c, w, r, tokenVerifier)
}

func processUpdateUser(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {
  var newUser models.User
  id := mux.Vars(r)["id"]
  tokenEmail, err := verifier.VerifyToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusForbidden)
    return
  }

  err = helpers.ReadJson(r.Body, &newUser)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  myKey, currentUserData, err := getUserWithId(c, id)

  if err != nil {
    var statusCode int
    if myKey == nil {
      statusCode = http.StatusNotFound
    } else {
      statusCode = http.StatusInternalServerError
    }

    helpers.SendError(w, err.Error(), statusCode)
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
  helpers.SendError(w, "", http.StatusNotImplemented)
}

func updateUserGroups(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  tokenVerifier := helpers.GetTokenVerifier(r)
  processUpdateUserGroups(c, w, r, tokenVerifier)
}

func processUpdateUserGroups(c appengine.Context, w http.ResponseWriter, r *http.Request, verifier helpers.TokenVerifier) {
  id := mux.Vars(r)["id"]
  _, err := verifier.VerifyToken(c, r)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusForbidden)
    return
  }

  var content updateUserGroupsContent
  err = helpers.ReadJson(r.Body, &content)
  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if content.GroupName == "" {
    helpers.SendError(w, "missing group name", http.StatusBadRequest)
    return
  }

  myKey, currentUserData, err := getUserWithId(c, id)

  if err != nil {
    var statusCode int
    if myKey == nil {
      statusCode = http.StatusNotFound
    } else {
      statusCode = http.StatusInternalServerError
    }

    helpers.SendError(w, err.Error(), statusCode)
    return
  }

  action := strings.ToLower(content.Action)
  if action == "remove" {
    err = removeGroupFromUser(c, myKey, currentUserData, content.GroupName)
  } else if action == "add" {
    err = addGroupFromUser(c, myKey, currentUserData, content.GroupName)
  } else {
    helpers.SendError(w, "invalid action", http.StatusBadRequest)
    return
  }

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = helpers.SendJson(w, currentUserData)

  if err != nil {
    helpers.SendError(w, err.Error(), http.StatusInternalServerError)
  }
}

func removeGroupFromUser(c appengine.Context, key *datastore.Key, currentUserData *models.User, groupName string) error {
  groupIndex := -1
  for i, a := range currentUserData.Groups {
    if a == groupName {
      c.Infof("group name index: %v", i)
      groupIndex = i
      break
    }
  }

  if groupIndex != -1 {
    currentUserData.Groups = append(currentUserData.Groups[:groupIndex], currentUserData.Groups[groupIndex+1:]...)
    c.Infof("groups to save: %v", currentUserData.Groups)
    _, err := datastore.Put(c, key, currentUserData)
    currentUserData.Id = key.Encode()

    if err != nil {
      return err
    }
  }
  return nil
}

func addGroupFromUser(c appengine.Context, key *datastore.Key, currentUserData *models.User, groupName string) error {
  groupAlreadyAdded := false
  for _, a := range currentUserData.Groups {
    if a == groupName {
      groupAlreadyAdded = true
      break
    }
  }

  if groupAlreadyAdded == false {
    currentUserData.Groups = append(currentUserData.Groups, groupName)

    _, err := datastore.Put(c, key, currentUserData)
    currentUserData.Id = key.Encode()

    if err != nil {
      return err
    }
  }
  return nil
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

func getUserWithId(c appengine.Context, id string) (*datastore.Key, *models.User, error) {
  myKey, err := datastore.DecodeKey(id)
  if err != nil {
    return nil, nil, err
  }

  var currentUserData models.User
  err = datastore.Get(c, myKey, &currentUserData)
  if err != nil {
    return myKey, nil, err
  }

  return myKey, &currentUserData, nil
}

type updateUserGroupsContent struct {
  Action string
  GroupName string
}
