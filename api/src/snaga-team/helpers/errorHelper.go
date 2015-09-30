package helpers

import (
	"net/http"
)

func SendError(w http.ResponseWriter, errorString string, code int) {
	http.Error(w, errorString, code)

	//TODO: Log this somewhere useful...
}
