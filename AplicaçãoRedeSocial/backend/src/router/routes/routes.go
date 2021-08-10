package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	isAuth   bool
}

func Config(r *mux.Router) {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
}
