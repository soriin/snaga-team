package controllers

import (
	"testing"
	// "fmt"
	"strings"
	"net/http"
	"bytes"

	"snaga-team/models"
	"snaga-team/test"
	"snaga-team/helpers"

	"appengine/datastore"
	"appengine/aetest"
)

func TestProcessAllUsers(t *testing.T) {
	c, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	w := test.NewFakeResponseWriter()
	want := "\"DisplayName\":\"snagaMan\",\"InGameName\":\"\""

	aUser := models.User{DisplayName: "snagaMan"}
	datastore.Put(c, datastore.NewIncompleteKey(c, "user", nil), &aUser)

	processAllUsers(c, w, nil)
	stringOutput := string(w.GetOutput()[:])

	if w.Calls["Write"] != 1 {
		t.Errorf("Expected Calls to FakeResponseWriter.Write to be 1 but was %v", w.Calls["Write"])
	}

	if strings.Contains(stringOutput, want) == false {
		t.Errorf("Expected output to contain %v but output was %v", want, stringOutput)
	}
}

func TestProcessAddUser(t *testing.T) {
	c, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Error(err)
	}
	defer c.Close()
	fakeVerifier := &test.FakeTokenVerifier{Email: "theman@getmoney.org"}
	want := "\"Email\":\"theman@getmoney.org\""
	w := test.NewFakeResponseWriter()
	user := models.User{DisplayName: "post test", FirstName: "Aegis", Email: "theman@getmoney.org"}
	var body bytes.Buffer
	err = helpers.SendJson(&body, &user)
	if err != nil {
		t.Error(err)
	}

	r, err := http.NewRequest("POST", "", &body)
	if err != nil {
		t.Error(err)
	}
	processAddUser(c, w, r, fakeVerifier)
	stringOutput := string(w.GetOutput()[:])

	if w.Calls["Write"] != 1 {
		t.Errorf("Expected Calls to FakeResponseWriter.Write to be 1 but was %v", w.Calls["Write"])
	}

	if strings.Contains(stringOutput, want) == false {
		t.Errorf("Expected output to contain %v but output was %v", want, stringOutput)
	}
}

func TestGetUserWithEmail(t *testing.T) {
	c, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Error(err)
	}
	defer c.Close()
  email := "theman@getmoney.org"

	aUser := models.User{DisplayName: "snagaMan", Email: email}
	datastore.Put(c, datastore.NewIncompleteKey(c, "user", nil), &aUser)

	user, err := getUserWithEmail(c, email)

	if err != nil {
		t.Fatalf("Unexpected error occured: %v", err.Error())
	}

	if user == nil {
		t.Errorf("Expected 1 user to be found but got none.")
	}
}
