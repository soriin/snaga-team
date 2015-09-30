package helpers

import (
	"fmt"
	"appengine"
	"net/http"
	appeng "google.golang.org/appengine"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/oauth2/v2"
)

func VerifyGoogleToken(c appengine.Context, r *http.Request) (string, error) {
	tokenHeader := r.Header["Auth-Token"]

	if len(tokenHeader) != 1 {
		return "", fmt.Errorf("No AuthToken in header")
	}
	tokenId := tokenHeader[0]

	ctx := appeng.NewContext(r)
	client, err := google.DefaultClient(ctx, oauth2.UserinfoProfileScope)
	if err != nil {
	  return "", err
	}

	oauth2Service, err := oauth2.New(client)
	if err != nil {
	  return "", err
	}

	tokenSvc := oauth2Service.Tokeninfo()
	tokenSvc = tokenSvc.IdToken(tokenId)
	tokenInfo, err := tokenSvc.Do()
	if err != nil {
	  return "", err
	}

	if tokenInfo.ExpiresIn == 0 {
	  return "", fmt.Errorf("Token expired!")
	}

	return tokenInfo.Email, nil;
}
