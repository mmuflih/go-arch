package p404

import (
	"encoding/json"
	"net/http"
)

type BaseHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type baseH struct{}

func NewBaseHandler() BaseHandler {
	return &baseH{}
}

func (this baseH) Handle(w http.ResponseWriter, r *http.Request) {
	data := "Sssssst! Silence is golden..."
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(
		data,
	)
}
