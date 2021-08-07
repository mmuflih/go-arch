package handlers

import (
	"net/http"

	"github.com/mmuflih/go-arch/context/user"
	"github.com/mmuflih/go-arch/http/core/request"
	"github.com/mmuflih/go-arch/http/requests"
	"github.com/mmuflih/golib/response"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:58:55
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
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

func (bh *authH) Register(w http.ResponseWriter, r *http.Request) {
	req := requests.RegisterRequest{}
	err := request.RequestValidator(r, bh.rr, &req)
	if err != nil {
		response.Exception(w, err, 422)
		return
	}
	err, resp := bh.handle.Register(req)
	response.Json(w, resp, err)
}

func (bh *authH) Login(w http.ResponseWriter, r *http.Request) {
	req := requests.LoginRequest{}
	err := request.RequestValidator(r, bh.rr, &req)
	if err != nil {
		response.Exception(w, err, 422)
		return
	}
	err, resp := bh.handle.Login(req)
	response.Json(w, resp, err)
}

func (bh *authH) Me(w http.ResponseWriter, r *http.Request) {
	userID := bh.auth.GetUserID(r)
	err, resp := bh.read.Me(userID)
	response.Json(w, resp, err)
}
