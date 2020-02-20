package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/routes/api"
)

type route struct {
	name       string
	method     string
	pattern    string
	HandleFunc http.HandlerFunc
}

type routes []route

/*Routes algo*/
type Routes interface {
	PointerToRouter()
}

/*Router algo*/
type Router struct {
	router *mux.Router
}

func generateARoutes(apiVar *api.API) []route {
	return routes{
		route{
			"showCarsList",
			http.MethodGet,
			"/cars",
			apiVar.GetListCars,
		},
		route{
			"getSpecificCar",
			http.MethodGet,
			"/cars/{carId}",
			apiVar.GetSpecificCar,
		},
		route{
			"addNewCar",
			http.MethodPost,
			"/cars/{carId}",
			apiVar.CreateNewCar,
		},
	}

}

/*CreateNewRInstance create a new router instance*/
func CreateNewRInstance() *Router {
	r := &Router{}
	api := api.CreateNewAPIInstance()

	r.router = mux.NewRouter().StrictSlash(true)

	routesVar := generateARoutes(api)

	for _, route := range routesVar {
		r.router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.HandleFunc)
	}

	return r
}

/*PointerToRouter return pointer to gorilla mux router, but is nil if do not run first run NewRouter()*/
func (r *Router) PointerToRouter() *mux.Router {
	return r.router
}
