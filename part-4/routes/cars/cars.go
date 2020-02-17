package cars

import (
	"fmt"
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/routes/cars/car"
)

/*Database All functions used for difference between users*/
type Database interface {
}

/*Cars All cars from data base*/
type Cars struct {
	cars map[string]car.Car
}

/*CreateNewInstance Create New Car instance*/
func CreateNewInstance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*GetListCars Create New Car instance*/
func GetListCars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*GetSpecificCar Create New Car instance*/
func GetSpecificCar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*AddItem Add new item in to the car*/
func AddItem(w http.ResponseWriter, r *http.Request) {
	car.AddItem(w, r)
}

/*ReturnListOfItems Return list with all elements in the car*/
func ReturnListOfItems(w http.ResponseWriter, r *http.Request) {
	car.ReturnListOfItems(w, r)
}

/*ReturnAItem Return a single item from the car, if not exist return a error*/
func ReturnAItem(w http.ResponseWriter, r *http.Request) {
	car.ReturnAItem(w, r)
}

/*UpdateAItem Change item from the car*/
func UpdateAItem(w http.ResponseWriter, r *http.Request) {
	car.UpdateAItem(w, r)
}

/*DeleteAItem Delete item from the car*/
func DeleteAItem(w http.ResponseWriter, r *http.Request) {
	car.DeleteAItem(w, r)
}
