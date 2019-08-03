package container

import (
	"github.com/mmuflih/go-di-arch/container/provider"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-05 10:10
**/

func BuildContainer() *dig.Container {
	c := dig.New()

	c = provider.BuildConfigProvider(c)
	c = provider.BuildRepositoryProvider(c)
	c = provider.BuildUseCaseProvider(c)
	c = provider.BuildHandlerProvider(c)

	return c
}
