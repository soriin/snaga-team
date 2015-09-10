package helpers

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

func SendJson(w io.Writer, obj interface{}) error {
	jStream, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jStream[:]))
	return nil
}

func ReadJson(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&x)
	return err
}
