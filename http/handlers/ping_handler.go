package handlers

import (
	"net/http"

	"github.com/mmuflih/go-arch/context/ping"
	"github.com/mmuflih/go-arch/http/requests"
	"github.com/mmuflih/golib/response"
)

type PingHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type pingH struct {
	puc ping.PingUsecase
}

func NewPingHandler(puc ping.PingUsecase) PingHandler {
	return &pingH{puc}
}

func (ph *pingH) Handle(w http.ResponseWriter, r *http.Request) {
	req := requests.PingRequest{}
	resp, err := ph.puc.Ping(req)
	response.Json(w, resp, err)
}
