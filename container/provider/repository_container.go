package provider

import (
	"github.com/mmuflih/go-di-arch/domain/repository/mysql"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildRepositoryProvider(c *dig.Container) *dig.Container {
	if err := c.Provide(mysql.NewUserRepo); err != nil {
		panic(err)
	}
	return c
}
