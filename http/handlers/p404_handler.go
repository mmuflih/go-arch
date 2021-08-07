package handlers

import (
	"net/http"

	"github.com/mmuflih/golib/response"
)

type P404Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type p404H struct{}

func NewP404Handler() P404Handler {
	return &p404H{}
}

func (bh p404H) Handle(w http.ResponseWriter, r *http.Request) {
	data := "Sssssst! Silence is golden..."
	response.Json(w, data, nil)
}
