package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func ReadParam(r *http.Request) (int64, error) {
	idString := chi.URLParam(r, "id")
	if idString == "" {
		return 0, errors.New("invalid ID parameter")
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return 0, errors.New("invalid ID parameter type")
	}
	return id, nil
}
