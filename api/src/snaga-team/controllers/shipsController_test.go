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

func TestProcessAllShips(t *testing.T) {
	c, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	w := test.NewFakeResponseWriter()
	want := "\"DisplayName\":\"test ship\",\"Manufacturer\":\"\""

	aShip := models.Ship{DisplayName: "test ship"}
	datastore.Put(c, datastore.NewIncompleteKey(c, "ship", nil), &aShip)

	processAllShips(c, w, nil)
	stringOutput := string(w.GetOutput()[:])

	if w.Calls["Write"] != 1 {
		t.Errorf("Expected Calls to FakeResponseWriter.Write to be 1 but was %v", w.Calls["Write"])
	}

	if strings.Contains(stringOutput, want) == false {
		t.Errorf("Expected output to contain %v but output was %v", want, stringOutput)
	}
}

func TestProcessAddShip(t *testing.T) {
	c, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Error(err)
	}
	defer c.Close()
	want := "\"DisplayName\":\"post test\",\"Manufacturer\":\"Aegis\""
	w := test.NewFakeResponseWriter()
	ship := models.Ship{DisplayName: "post test", Manufacturer: "Aegis"}
	var body bytes.Buffer
	err = helpers.SendJson(&body, &ship)
	if err != nil {
		t.Error(err)
	}

	r, err := http.NewRequest("POST", "", &body)
	if err != nil {
		t.Error(err)
	}
	processAddShip(c, w, r)
	stringOutput := string(w.GetOutput()[:])

	if w.Calls["Write"] != 1 {
		t.Errorf("Expected Calls to FakeResponseWriter.Write to be 1 but was %v", w.Calls["Write"])
	}

	if strings.Contains(stringOutput, want) == false {
		t.Errorf("Expected output to contain %v but output was %v", want, stringOutput)
	}
}
