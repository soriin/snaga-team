package helpers

import (
	"encoding/json"
	"io"
	"fmt"
)

func SendJson(w io.Writer, obj interface{}) error {
	jStream, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jStream[:]))
	return nil
}
