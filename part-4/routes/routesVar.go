package routes

import "net/http"

type route struct {
	name       string
	method     string
	pattern    string
	HandleFunc http.HandlerFunc
}

type routes []route

var routesVar = routes{
	route{
		"showCarsList",
		"GET",
		"/cars",
		Index,
	},
	route{
		"getSpecificCar",
		"GET",
		"/cars/{carId}",
		Index,
	},
	route{
		"addNewCar",
		"Post",
		"/cars/{carId}",
		Index,
	},
	route{
		"returnSpecificItemOfCar",
		"GET",
		"/cars/{carId}/{itemId}",
		Index,
	},
	route{
		"addNewItemToCar",
		"POST",
		"/cars/{carId}/{itemId}",
		Index,
	},
	route{
		"updateElementFromACar",
		"PUT",
		"/cars/{carId}/{itemId}",
		Index,
	},
	route{
		"removeElementFromACar",
		"DELETE",
		"/cars/{carId}/{itemId}",
		Index,
	},
}
