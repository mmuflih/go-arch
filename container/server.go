package container

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-di-arch/role"
	"github.com/mmuflih/golib/middleware"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-05 10:59
**/

type ServerRoute struct {
	config  conf.Config
	handler http.Handler
	router  *mux.Router
}

func NewRoute(c conf.Config, handler http.Handler, router *mux.Router) *ServerRoute {
	myrole := make(map[string][]string)

	myrole[role.ADMIN] = []string{role.ADMIN}
	myrole[role.LEADER] = []string{role.LEADER, role.ADMIN}
	myrole[role.USER] = []string{role.USER, role.LEADER, role.ADMIN}

	middleware.InitJWTMiddlewareWithRole([]byte(c.GetString("key")), jwt.SigningMethodHS512, myrole)

	for _, l := range c.GetStringSlice("env_label") {
		fmt.Println(l)
	}

	return &ServerRoute{c, handler, router}
}

func (s *ServerRoute) Run() {
	log.Println("Application is running at ", time.Now().Format("2006-01-02 15:04:05.000"))
	log.Println("Server listen on", s.config.GetString(`server.address`))
	log.Fatal(http.ListenAndServe(s.config.GetString(`server.address`), s.handler))
}
