package provider

import "github.com/mmuflih/go-arch/http/handlers"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func Handlers() []interface{} {
	var h []interface{}

	h = append(h, handlers.NewPingHandler)
	h = append(h, handlers.NewP404Handler)
	h = append(h, handlers.NewAuthHandler)

	return h
}
