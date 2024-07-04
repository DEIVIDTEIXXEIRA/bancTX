package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri      string
	Metodo   string
	function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Config(r *mux.Router) *mux.Router {
	routes := routeFromUser
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.function).Methods(route.Metodo)
	}
	return r
}
