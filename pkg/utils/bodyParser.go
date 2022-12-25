package utils

import (
	"encoding/json"
	"io"
)

func BodyParser(body io.ReadCloser, model interface{}) error {
	if err := json.NewDecoder(body).Decode(model); err != nil {
		return err
	}

	return nil
}
