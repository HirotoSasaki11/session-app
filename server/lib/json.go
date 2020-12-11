package lib

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, v interface{}) error {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func BodyToJson(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
