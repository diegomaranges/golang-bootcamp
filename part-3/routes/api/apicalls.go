package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/dbcars"
)

/*Interactions algo*/
type Interactions interface {
	GetListCars(w http.ResponseWriter, r *http.Request)
	GetSpecificCar(w http.ResponseWriter, r *http.Request)
	CreateNewCar(w http.ResponseWriter, r *http.Request)
}

/*API algo*/
type API struct {
	cars *dbcars.Cars
}

func generateAResponse(v interface{}) {
	fmt.Print(v)
}

/*CreateNewAPIInstance algo*/
func CreateNewAPIInstance() *API {
	apiVar := &API{}
	var err error
	apiVar.cars, err = dbcars.CreateNewCarsInstance()
	if err != nil {
		return nil
	}
	return apiVar
}

/*GetListCars algo*/
func (a *API) GetListCars(w http.ResponseWriter, r *http.Request) {

}

/*GetSpecificCar algo*/
func (a *API) GetSpecificCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carID := params["carId"]

	data, err := a.cars.ReturnASpecificCar(carID)
	if err != nil {
		generateAResponse(err)
	}
	generateAResponse(data)
}

/*CreateNewCar algo*/
func (a *API) CreateNewCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carID := params["carId"]

	err := a.cars.CreateNewCar(carID)
	generateAResponse(err)

}
