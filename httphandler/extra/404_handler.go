package extra

import (
	"encoding/json"
	"net/http"
)

type P404Handler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type p404handler struct{}

func NewP404Handler() P404Handler {
	return &p404handler{}
}

func (this p404handler) Handle(w http.ResponseWriter, r *http.Request) {
	data := "Sssssst! Silence is golden..."
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(
		data,
	)
}
