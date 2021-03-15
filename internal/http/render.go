package http

import (
	"encoding/json"
	"io"
	"net/http"
)

// RenderJSON writes an http response using the value passed in v as JSON.
// If it cannot convert the value to JSON, it returns an error.
func RenderJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	js, err := json.Marshal(v)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(js)

	if err != nil {
		return err
	}

	return nil
}

// DecodeJSON is a helper that abstracts the decoding of JSON requests.
func DecodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}