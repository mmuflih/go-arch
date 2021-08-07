package main

import (
	"fmt"

	"github.com/mmuflih/go-arch/container"
	"go.uber.org/dig"
)

var _ = dig.Name

func main() {
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
