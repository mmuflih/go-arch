package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/go-arch/context/user"
	"github.com/mmuflih/go-arch/http/core/request"
	"github.com/mmuflih/go-arch/http/core/response"
	"github.com/mmuflih/go-arch/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:58:55
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type BaseHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Me(c *gin.Context)
}

type authH struct {
	handle user.Handler
	read   user.Reader
	auth   user.GetAuthUserUsecase
	rr     request.Reader
}

func NewAuthHandler(handle user.Handler, read user.Reader,
	auth user.GetAuthUserUsecase,
	rr request.Reader) AuthHandler {
	return &authH{handle, read, auth, rr}
}

func (bh *baseHandler) Register(c *gin.Context) {
	req := requests.RegisterRequest{}
	err := request.RequestValidator(c.Request, bh.rr, &req)
	if err != nil {
		response.Exception(c, err, 422)
		return
	}
	err, resp := bh.handle.Register(req)
	response.Json(c, resp, err)
}

func (bh *baseHandler) Login(c *gin.Context) {
	req := requests.LoginRequest{}
	err := request.RequestValidator(c.Request, bh.rr, &req)
	if err != nil {
		response.Exception(c, err, 422)
		return
	}
	err, resp := bh.handle.Login(req)
	response.Json(c, resp, err)
}

func (bh *baseHandler) Me(c *gin.Context) {
	userID := bh.auth.GetUserID(c.Request)
	err, resp := bh.read.Me(userID)
	response.Json(c, resp, err)
}
