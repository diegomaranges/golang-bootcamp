package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func generateAResponse() {

}

/*GetListCars algo*/
func GetListCars(w http.ResponseWriter, r *http.Request) {

}

/*GetSpecificCar algo*/
func GetSpecificCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID := params["carId"]

}

/*CreateNewCar algo*/
func CreateNewCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID := params["carId"]

}
