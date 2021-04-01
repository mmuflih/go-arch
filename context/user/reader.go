package user

import (
	"github.com/mmuflih/go-di-arch/domain/repository/mysql"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-29 01:54:09
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type Reader interface {
	Me(id uint64) (error, interface{})
}

type read struct {
	uRepo mysql.UserRepository
}

func NewReader(
	uRepo mysql.UserRepository,
) Reader {
	return &read{uRepo}
}
