package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-arch/config"
	"github.com/mmuflih/go-arch/http/core/request"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildConfigProvider(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(func(c conf.Config) *config.Keys {
		signature := []byte(c.GetString("key"))
		return &config.Keys{
			Signature: signature,
		}
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func() conf.Config {
		return conf.NewConfig()
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func(c conf.Config) (*config.MongoDB, error) {
		return config.NewMongoDBConnections(c)
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func(c conf.Config) (*config.MySQL, error) {
		return config.NewMySQLConnections(c)
	}); err != nil {
		panic(err)
	}

	err = c.Provide(func() request.Reader {
		return request.NewRequestReader()
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func() *gin.Engine {
		return gin.Default()
	})
	if err != nil {
		panic(err)
	}

	return c
}
