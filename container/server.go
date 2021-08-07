package container

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-arch/role"
	"github.com/mmuflih/golib/middleware"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-05 10:59
**/

type ServerRoute struct {
	config conf.Config
	router *gin.Engine
}

func NewRoute(c conf.Config, router *gin.Engine) *ServerRoute {
	myrole := make(map[string][]string)

	myrole[role.ADMIN] = []string{role.ADMIN}
	myrole[role.LEADER] = []string{role.LEADER, role.ADMIN}
	myrole[role.USER] = []string{role.USER, role.LEADER, role.ADMIN}

	middleware.InitJWTMiddlewareWithRole([]byte(c.GetString("key")), jwt.SigningMethodHS512, myrole)

	for _, l := range c.GetStringSlice("env_label") {
		fmt.Println(l)
	}

	/** init cors */
	router.Use(cors.New(cors.Config{
		AllowOrigins:     c.GetStringSlice("cors.allowed_origins"),
		AllowMethods:     c.GetStringSlice("cors.allowed_methods"),
		AllowHeaders:     c.GetStringSlice("cors.allowed_headers"),
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	return &ServerRoute{c, router}
}

func (s *ServerRoute) Run() {
	log.Println("Application is running at ", time.Now().Format("2006-01-02 15:04:05.000"))
	port := s.config.GetString(`server.address`)
	gin.SetMode(s.config.GetString("env"))
	s.router.Run(port)
}
