package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/go-arch/context/ping"
	"github.com/mmuflih/go-arch/http/core/response"
	"github.com/mmuflih/go-arch/http/requests"
)

type PingHandler interface {
	Handle(c *gin.Context)
}

type pingH struct {
	puc ping.PingUsecase
}

func NewPingHandler(puc ping.PingUsecase) PingHandler {
	return &pingH{puc}
}

func (ph *pingH) Handle(c *gin.Context) {
	req := requests.PingRequest{}
	resp, err := ph.puc.Ping(req)
	response.Json(c, resp, err)
}
