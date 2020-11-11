package user

import (
	"net/http"

	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/context/user"
	"github.com/mmuflih/go-httplib/httplib"
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
	registerUC user.RegisterUsecase
	loginUC    user.LoginUsecase
	meUC       user.GetMeUsecase
	auth       user.GetAuthUserUsecase
	rr         httplib.RequestReader
}

func NewBaseHandler(registerUC user.RegisterUsecase, loginUC user.LoginUsecase,
	rr httplib.RequestReader, meUC user.GetMeUsecase,
	auth user.GetAuthUserUsecase) BaseHandler {
	return &baseHandler{registerUC, loginUC, meUC, auth, rr}
}

func (bh *baseHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := registerRequest{}
	err := app.RequestValidator(r, bh.rr, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := bh.registerUC.Handle(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	httplib.ResponseData(w, resp)
	return
}

func (bh *baseHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	err := app.RequestValidator(r, bh.rr, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := bh.loginUC.Handle(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	httplib.ResponseData(w, resp)
	return
}

func (bh *baseHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := bh.auth.GetUserID(r)
	err, resp := bh.meUC.Me(userID)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	httplib.ResponseData(w, resp)
	return
}
