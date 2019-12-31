package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mmuflih/go-di-arch/container"
	"github.com/mmuflih/go-di-arch/role"
	"github.com/mmuflih/go-httplib/httplib"
)

import "go.uber.org/dig"

var _ = dig.Name

func main() {
	myrole := make(map[string][]string)

	myrole[role.ADMIN] = []string{role.ADMIN}
	myrole[role.LEADER] = []string{role.LEADER, role.ADMIN}
	myrole[role.USER] = []string{role.USER, role.LEADER, role.ADMIN}

	httplib.InitJWTMiddlewareWithRole([]byte("Go-DI-arch"), jwt.SigningMethodHS512, myrole)

	c := container.BuildContainer()

	if err := c.Invoke(container.InvokeRoute); err != nil {
		panic(err)
	}

	if err := c.Provide(container.NewRoute); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(s *container.ServerRoute) {
		s.Run()
	}); err != nil {
		fmt.Println(err)
	}
}
