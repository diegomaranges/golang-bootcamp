package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/router/apifunctions"
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
			"showCarsList",
			http.MethodGet,
			"/cars/{carID}",
			apifunctions.ReturnCar,
		},
		route{
			"getSpecificCar",
			http.MethodGet,
			"/cars/{carID}/{itemID}",
			apifunctions.ReturnItem,
		},
		route{
			"addNewCar",
			http.MethodPost,
			"/cars/{carID}/{itemID}",
			apifunctions.AddItem,
		},
	}

}

/*NewRute create a new router*/
func NewRute() *mux.Router {
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
