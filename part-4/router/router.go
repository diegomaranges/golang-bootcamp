package router

import (
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/router/apifunctions"

	"github.com/gorilla/mux"
)

type route struct {
	name       string
	method     string
	pattern    string
	HandleFunc http.HandlerFunc
}

type routes []route

func generateARoutes() []route {
	return routes{
		route{
			"SignIn",
			http.MethodPost,
			"/signin",
			apifunctions.Signin,
		},
		route{
			"Refresh",
			http.MethodPost,
			"/refresh",
			apifunctions.Refresh,
		}, /*
			route{
				"CreateNewCar",
				http.MethodPost,
				"/cars/{carID}",
				apifunctions.CreateNewCar,
			},*/
		route{
			"getSpecificCar",
			http.MethodGet,
			"/cars/{carID}",
			apifunctions.ReturnCar,
		}, /*
			route{
				"DeleteCar",
				http.MethodDelete,
				"/cars/{carID}",
				apifunctions.DeleteCar,
			},*/
		route{
			"getSpecificItem",
			http.MethodGet,
			"/cars/{carID}/{itemID}",
			apifunctions.ReturnItem,
		},
		route{
			"addItem",
			http.MethodPost,
			"/cars/{carID}/{itemID}",
			apifunctions.AddItem,
		},
		route{
			"updateItem",
			http.MethodPut,
			"/cars/{carID}/{itemID}",
			apifunctions.UpdateItem,
		},
		route{
			"deleteItem",
			http.MethodDelete,
			"/cars/{carID}/{itemID}",
			apifunctions.DeleteItem,
		},
	}

}

/*NewRoute create a new router*/
func NewRoute() *mux.Router {
	newRoute := mux.NewRouter().StrictSlash(true)

	for _, route := range generateARoutes() {
		newRoute.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.HandleFunc)
	}

	return newRoute
}
