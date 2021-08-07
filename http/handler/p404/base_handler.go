package p404

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler interface {
	Handle(c *gin.Context)
}

type baseH struct{}

func NewBaseHandler() BaseHandler {
	return &baseH{}
}

func (bh baseH) Handle(c *gin.Context) {
	data := "Sssssst! Silence is golden..."
	c.JSONP(http.StatusOK, data)
}
