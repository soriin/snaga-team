package test

import (
	"appengine"
	"net/http"
)

type FakeTokenVerifier struct {
	Email string
	Error error
}

func (v *FakeTokenVerifier) VerifyToken(appengine.Context, *http.Request) (string, error) {
	if v.Error == nil {
		return v.Email, nil
	}
	return "", v.Error
}
