package ping

import (
	"net/http"

	"github.com/mmuflih/go-di-arch/context/ping"
	"github.com/mmuflih/go-di-arch/http/requests"
	"github.com/mmuflih/golib/response"
)

type BaseHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type baseH struct {
	puc ping.PingUsecase
}

func NewBaseHandler(puc ping.PingUsecase) BaseHandler {
	return &baseH{puc}
}

func (this *baseH) Handle(w http.ResponseWriter, r *http.Request) {
	req := requests.PingRequest{}
	resp, err := this.puc.Ping(req)
	response.Json(w, resp, err)
}
