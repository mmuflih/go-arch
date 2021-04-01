package auth

import (
	"net/http"

	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/context/user"
	"github.com/mmuflih/go-di-arch/http/requests"
	"github.com/mmuflih/golib/request"
	"github.com/mmuflih/golib/response"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:58:55
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type BaseHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
}

type baseHandler struct {
	handle user.Handler
	read   user.Reader
	auth   user.GetAuthUserUsecase
	rr     request.Reader
}

func NewBaseHandler(handle user.Handler, read user.Reader,
	auth user.GetAuthUserUsecase,
	rr request.Reader) BaseHandler {
	return &baseHandler{handle, read, auth, rr}
}

func (bh *baseHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := requests.RegisterRequest{}
	err := app.RequestValidator(r, bh.rr, &req)
	if err != nil {
		response.Exception(w, err, 422)
		return
	}
	err, resp := bh.handle.Register(req)
	response.Json(w, resp, err)
}

func (bh *baseHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := requests.LoginRequest{}
	err := app.RequestValidator(r, bh.rr, &req)
	if err != nil {
		response.Exception(w, err, 422)
		return
	}
	err, resp := bh.handle.Login(req)
	response.Json(w, resp, err)
}

func (bh *baseHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := bh.auth.GetUserID(r)
	err, resp := bh.read.Me(userID)
	response.Json(w, resp, err)
}
