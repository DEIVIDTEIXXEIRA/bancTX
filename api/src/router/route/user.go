package route

import (
	"net/http"

	"github.com/bankTX/api/src/controllers"
)

var routeFromUser = []Route{
	{
		Uri:      "/user",
		Metodo:   http.MethodPost,
		function: controllers.InsertUser,
		Auth:     false,
	},
	{
		Uri:      "/user",
		Metodo:   http.MethodGet,
		function: controllers.GetUser,
		Auth:     false,
	},
	{
		Uri:      "/user/{userID}",
		Metodo:   http.MethodPut,
		function: controllers.UpdateUser,
		Auth:     false,
	},
	{
		Uri:      "/user/{userID}",
		Metodo:   http.MethodDelete,
		function: controllers.DeleteUser,
		Auth:     false,
	},
}
