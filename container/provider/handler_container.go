package provider

import (
	"github.com/mmuflih/go-di-arch/http/handler/auth"
	"github.com/mmuflih/go-di-arch/http/handler/p404"
	"github.com/mmuflih/go-di-arch/http/handler/ping"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func Handlers() []interface{} {
	var h []interface{}

	h = append(h, p404.NewBaseHandler)
	h = append(h, ping.NewBaseHandler)
	h = append(h, auth.NewBaseHandler)

	return h
}
