package cars

import (
	"fmt"
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/routes/cars/car"
)

/*Database All functions used for difference between users
type Database interface {
	GetListCars(w http.ResponseWriter, r *http.Request)
	GetSpecificCar(w http.ResponseWriter, r *http.Request)
	AddItem(w http.ResponseWriter, r *http.Request)
	ReturnListOfItems(w http.ResponseWriter, r *http.Request)
	ReturnAItem(w http.ResponseWriter, r *http.Request)
	UpdateAItem(w http.ResponseWriter, r *http.Request)
	DeleteAItem(w http.ResponseWriter, r *http.Request)
}*/

/*Cars All cars from data base*/
type Cars struct {
	cars map[string]*car.Car
}

var myVar *Cars

/*CreateNewCarsInstance Create New Car instance*/
func CreateNewCarsInstance() {
	newCars := &Cars{}
	newCars.cars = make(map[string]*car.Car)
	newCars.cars["1"] = car.CreateNewCarInstance()
	myVar = newCars
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
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	myVar.cars["1"].AddItem(w, r)
	/*id := "1"
	c.cars[id].AddItem(w, r)*/
}

/*ReturnListOfItems Return list with all elements in the car*/
func ReturnListOfItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	/*id := "1"
	c.cars[id].ReturnListOfItems(w, r)*/
}

/*ReturnAItem Return a single item from the car, if not exist return a error*/
func ReturnAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	/*id := "1"
	c.cars[id].ReturnAItem(w, r)*/
}

/*UpdateAItem Change item from the car*/
func UpdateAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	/*id := "1"
	c.cars[id].UpdateAItem(w, r)*/
}

/*DeleteAItem Delete item from the car*/
func DeleteAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	/*id := "1"
	c.cars[id].DeleteAItem(w, r)*/
}
