package container

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmuflih/go-arch/http/handlers"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *mux.Router,
	pingH handlers.PingHandler, p404H handlers.P404Handler,
	userH handlers.AuthHandler,
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
