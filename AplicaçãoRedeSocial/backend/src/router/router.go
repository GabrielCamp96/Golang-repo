package router

import (
	router "api/src/router/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	r := mux.NewRouter()
	router.Config(r)

	return r
}
