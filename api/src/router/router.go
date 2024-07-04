package router

import (
	"github.com/bankTX/api/src/router/route"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return route.Config(r)
}
