package container

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmuflih/go-di-arch/http/handler/auth"
	"github.com/mmuflih/go-di-arch/http/handler/p404"
	"github.com/mmuflih/go-di-arch/http/handler/ping"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *mux.Router,
	pingH ping.BaseHandler, p404H p404.BaseHandler, userH auth.BaseHandler,
) {
	route.NotFoundHandler = http.HandlerFunc(p404H.Handle)
	/** api v1 route */
	apiV1 := route.PathPrefix("/api/v1").Subrouter()
	pingRoute := apiV1.PathPrefix("/ping").Subrouter()
	userRoute := apiV1.PathPrefix("/user").Subrouter()

	/** ping */
	pingRoute.HandleFunc("", pingH.Handle).Methods("GET")

	/** user */
	userRoute.HandleFunc("/register", userH.Register).Methods("POST")
	userRoute.HandleFunc("/login", userH.Login).Methods("POST")
	userRoute.HandleFunc("", userH.Me).Methods("GET")
}
