package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*Routes algo*/
type Routes interface {
	NewRouter()
	PointerToRouter()
}

/*Router algo*/
type Router struct {
	router *mux.Router
}

/*NewRouter algo*/
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

/*PointerToRouter asd*/
func (r *Router) PointerToRouter() *mux.Router {
	return r.router
}

/*Index algo*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*CreateNewRouter algo
func CreateNewRouter() *Router {
	var router *Router
	return router
}*/
