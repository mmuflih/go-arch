package provider

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-di-arch/config"
	"github.com/mmuflih/go-httplib/httplib"
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

	if err = c.Provide(func(c conf.Config) (*config.MongoConnections, error) {
		return config.NewMongoDB2(c)
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func(c conf.Config) (*config.MyConn, error) {
		return config.NewMysqlGormConn(c)
	}); err != nil {
		panic(err)
	}

	err = c.Provide(func() httplib.RequestReader {
		return httplib.NewMuxRequestReader()
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func() *mux.Router {
		return mux.NewRouter()
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(c conf.Config, api *mux.Router) http.Handler {
		return config.InitCors(c, api)
	})
	if err != nil {
		panic(err)
	}

	return c
}
