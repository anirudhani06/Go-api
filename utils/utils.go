package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anirudhani06/Go-api/helpers"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing Request Body")
	}

	return json.NewDecoder(r.Body).Decode(payload)

}

func WriteJson(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)

}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	response := helpers.Response{
		Status: false,
		Data:   map[string]string{},
		Errors: err.Error(),
	}
	WriteJson(w, statusCode, response)
}
