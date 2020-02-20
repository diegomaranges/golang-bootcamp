package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/api"
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
	CreateNewRInstance()
	NewRouter()
	PointerToRouter()
}

/*Router algo*/
type Router struct {
	router *mux.Router
}

var routesVar = routes{
	route{
		"showCarsList",
		http.MethodGet,
		"/cars",
		api.GetListCars,
	},
	route{
		"getSpecificCar",
		http.MethodGet,
		"/cars/{carId}",
		api.GetSpecificCar,
	},
	route{
		"addNewCar",
		http.MethodPost,
		"/cars/{carId}",
		api.CreateNewCar,
	},
}

/*CreateNewRInstance create a new router instance*/
func CreateNewRInstance() *Router {
	return &Router{}
}

/*NewRouter initializate router var with paths, methods and hangler function*/
func (r *Router) NewRouter() {
	r.router = mux.NewRouter().StrictSlash(true)

	for _, route := range routesVar {
		r.router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.HandleFunc)
	}
}

/*PointerToRouter return pointer to gorilla mux router, but is nil if do not run first run NewRouter()*/
func (r *Router) PointerToRouter() *mux.Router {
	return r.router
}
