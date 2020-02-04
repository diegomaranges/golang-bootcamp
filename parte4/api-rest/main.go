package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var routes = Routes {
	Route {
		"Index",
		"GET",
		"/",
		Index
	}
	Route {
		"Index",
		"GET",
		"/car",
		handleCar,
	}
	Route {
		"Index",
		"GET",
		"/car/{id}",
		carShow,
	}
}

/*Route algo*/
type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandleFunc
}

/*Routes*/
type Routes []Route

/*Car algo*/
type Car struct {
	Name               string `json:"name"`
	Year               int    `json:"year"`
	QuantityOfElements int    `json:"quantity"`
}

/*Cars algo*/
type Cars []Car

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi server")
}

func handleCar(w http.ResponseWriter, r *http.Request) {
	cars := Cars{
		Car{"1", 1, 1},
		Car{"2", 2, 12},
		Car{"3", 3, 123},
		Car{"4", 4, 1234},
	}

	json.NewEncoder(w).Encode(cars)
}

func carShow(w http.ResponseWriter, r *http.Request) {
	cars := Cars{
		Car{"1", 1, 1},
		Car{"2", 2, 12},
		Car{"3", 3, 123},
		Car{"4", 4, 1234},
	}
	params := mux.Vars(r)
	carID := params["id"]
	founded := false
	var id int

	for _, car := range cars {
		if car.Name == carID {
			founded = true
			id, _ = strconv.Atoi(car.Name)
			break
		}
	}

	if founded {
		json.NewEncoder(w).Encode(cars[id-1])
	}
	//fmt.Fprintf(w, "Has cargado la pelicula numero %s", car_id)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := routes {
		route.
			Name(route.Name).
			Method(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
}

