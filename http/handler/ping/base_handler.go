package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/go-arch/context/ping"
	"github.com/mmuflih/go-arch/http/core/response"
	"github.com/mmuflih/go-arch/http/requests"
)

type BaseHandler interface {
	Handle(c *gin.Context)
}

type baseH struct {
	puc ping.PingUsecase
}

func NewBaseHandler(puc ping.PingUsecase) BaseHandler {
	return &baseH{puc}
}

func (this *baseH) Handle(c *gin.Context) {
	req := requests.PingRequest{}
	resp, err := this.puc.Ping(req)
	response.Json(c, resp, err)
}
