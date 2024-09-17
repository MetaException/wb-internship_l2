package utils

import (
	"encoding/json"
	"io"
)

func Serialize(w io.Writer, data any) error {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
