package helpers

import (
	"net/http"
	"appengine"
)

type TokenVerifier interface {
	VerifyToken(appengine.Context, *http.Request) (string, error)
}

func GetTokenVerifier(r *http.Request) TokenVerifier {
	return &GoogleTokenVerifier{}
}
