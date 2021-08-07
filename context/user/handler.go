package user

import (
	"github.com/mmuflih/go-arch/config"
	"github.com/mmuflih/go-arch/domain/repository/mysql"
	"github.com/mmuflih/go-arch/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:46:02
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type Handler interface {
	Register(req requests.RegisterRequest) (error, interface{})
	Login(req requests.LoginRequest) (error, interface{})
}

type handle struct {
	uRepo  mysql.UserRepository
	ueRepo mysql.UserEmailRepository
	upRepo mysql.UserPasswordRepository
	keys   *config.Keys
}

func NewHandler(userRepo mysql.UserRepository,
	userEmailRepo mysql.UserEmailRepository,
	userPasswordRepo mysql.UserPasswordRepository,
	keys *config.Keys) Handler {
	return &handle{userRepo, userEmailRepo, userPasswordRepo, keys}
}
