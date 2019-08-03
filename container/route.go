package container

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmuflih/go-di-arch/httphandler/extra"
	"github.com/mmuflih/go-di-arch/httphandler/ping"
	"github.com/mmuflih/go-di-arch/httphandler/user"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *mux.Router,
	pingH ping.PingHandler, p404H extra.P404Handler, userAdd user.AddHandler,
	userEdit user.EditHandler, userGet user.GetHandler, userList user.ListHandler,
) {
	route.NotFoundHandler = http.HandlerFunc(p404H.Handle)
	/** api v1 route */
	apiV1 := route.PathPrefix("/api/v1").Subrouter()
	pingRoute := apiV1.PathPrefix("/ping").Subrouter()
	userRoute := apiV1.PathPrefix("/user").Subrouter()

	/** ping */
	pingRoute.HandleFunc("", pingH.Handle).Methods("GET")

	/** user */
	userRoute.HandleFunc("", userAdd.Handle).Methods("POST")
	userRoute.HandleFunc("", userList.Handle).Methods("GET")
	userRoute.HandleFunc("/{id}", userEdit.Handle).Methods("PUT")
	userRoute.HandleFunc("/{id}", userGet.Handle).Methods("GET")
}
