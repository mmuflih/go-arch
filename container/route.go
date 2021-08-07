package container

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/go-arch/http/handler/auth"
	"github.com/mmuflih/go-arch/http/handler/p404"
	"github.com/mmuflih/go-arch/http/handler/ping"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *gin.Engine,
	pingH ping.BaseHandler, p404H p404.BaseHandler, userH auth.BaseHandler,
) {
	route.NoRoute(p404H.Handle)
	/** api v1 route */
	apiV1 := route.Group("/api/v1")

	/** ping */
	pingRoute := apiV1.Group("/ping")
	pingRoute.GET("", pingH.Handle)

	/** user */
	userRoute := apiV1.Group("/user")
	userRoute.POST("/register", userH.Register)
	userRoute.POST("/login", userH.Login)
	userRoute.GET("", userH.Me)
}
