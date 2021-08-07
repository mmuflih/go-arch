package provider

import (
	"github.com/mmuflih/go-arch/context/ping"
	"github.com/mmuflih/go-arch/context/user"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func Usecases() []interface{} {

	var u []interface{}

	u = append(u, ping.NewPingUsecase)
	u = append(u, user.NewHandler)
	u = append(u, user.NewGetAuthUserUsecase)
	u = append(u, user.NewReader)

	return u
}
