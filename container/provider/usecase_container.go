package provider

import (
	"github.com/mmuflih/go-di-arch/context/ping"
	"github.com/mmuflih/go-di-arch/context/user"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildUseCaseProvider(c *dig.Container) *dig.Container {
	if err := c.Provide(ping.NewPingUsecase); err != nil {
		panic(err)
	}

	if err := c.Provide(user.NewGenerateUsecase); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewGetAuthUserUsecase); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewGetMeUsecase); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewLoginUsecase); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewRegisterUsecase); err != nil {
		panic(err)
	}

	return c
}
