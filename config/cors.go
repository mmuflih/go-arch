package config

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/golib/logger"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-05 10:54
**/

func InitCors(config conf.Config, api *mux.Router) http.Handler {
	headersVal := config.GetStringSlice("cors.allowed_headers")
	methodsVal := config.GetStringSlice("cors.allowed_methods")
	originsVal := config.GetStringSlice("cors.allowed_origins")
	headers := handlers.AllowedHeaders(headersVal)
	methods := handlers.AllowedMethods(methodsVal)
	origins := handlers.AllowedOrigins(originsVal)
	cors := handlers.CORS(headers, methods, origins)
	apiCors := cors(logger.Log(api))
	return apiCors
}
