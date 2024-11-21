package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// TO unmarshal json to struct

func ParseBody(r *http.Request, x interface{}) {
	if _, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte{}, x); err != nil {
			return
		}
	}
}