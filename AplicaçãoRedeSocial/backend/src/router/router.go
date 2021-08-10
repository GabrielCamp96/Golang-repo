package router

import (
	"api/src/middleware"
	router "api/src/router/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	r := mux.NewRouter()
	router.Config(r)
	r.Use(middleware.SetContentTypeApplicationJson)
	return r
}
