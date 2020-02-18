package routes

import (
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/routes/cars"
)

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
		cars.GetListCars,
	},
	route{
		"getSpecificCar",
		"GET",
		"/cars/{carId}",
		cars.GetSpecificCar,
	}, /*
		route{
			"addNewCar",
			"POST",
			"/cars/{carId}",
			cars.CreateNewInstance,
		},*/
	route{
		"returnListOfItem",
		"GET",
		"/cars/{carId}/{itemId}",
		cars.ReturnListOfItems,
	},
	route{
		"returnSpecificItemOfCar",
		"GET",
		"/cars/{carId}/{itemId}",
		cars.ReturnAItem,
	},
	route{
		"addNewItemToCar",
		"POST",
		"/cars/{carId}/{itemId}",
		cars.AddItem,
	},
	route{
		"updateElementFromACar",
		"PUT",
		"/cars/{carId}/{itemId}",
		cars.UpdateAItem,
	},
	route{
		"removeElementFromACar",
		"DELETE",
		"/cars/{carId}/{itemId}",
		cars.DeleteAItem,
	},
}
