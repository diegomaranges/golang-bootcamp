package cars

import (
	"fmt"
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/routes/cars/car"
)

/*Database All functions used for difference between users*/
type Database interface {
	GetListCars(w http.ResponseWriter, r *http.Request)
	GetSpecificCar(w http.ResponseWriter, r *http.Request)
	AddItem(w http.ResponseWriter, r *http.Request)
	ReturnListOfItems(w http.ResponseWriter, r *http.Request)
	ReturnAItem(w http.ResponseWriter, r *http.Request)
	UpdateAItem(w http.ResponseWriter, r *http.Request)
	DeleteAItem(w http.ResponseWriter, r *http.Request)
}

/*Cars All cars from data base*/
type Cars struct {
	cars map[string]*car.Car
}

/*CreateNewInstance Create New Car instance*/
func CreateNewInstance(w http.ResponseWriter, r *http.Request) *Cars {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	return &Cars{}
}

/*GetListCars Create New Car instance*/
func (c *Cars) GetListCars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*GetSpecificCar Create New Car instance*/
func (c *Cars) GetSpecificCar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*AddItem Add new item in to the car*/
func (c *Cars) AddItem(w http.ResponseWriter, r *http.Request) {
	id := "1"
	c.cars[id].AddItem(w, r)
}

/*ReturnListOfItems Return list with all elements in the car*/
func (c *Cars) ReturnListOfItems(w http.ResponseWriter, r *http.Request) {
	id := "1"
	c.cars[id].ReturnListOfItems(w, r)
}

/*ReturnAItem Return a single item from the car, if not exist return a error*/
func (c *Cars) ReturnAItem(w http.ResponseWriter, r *http.Request) {
	id := "1"
	c.cars[id].ReturnAItem(w, r)
}

/*UpdateAItem Change item from the car*/
func (c *Cars) UpdateAItem(w http.ResponseWriter, r *http.Request) {
	id := "1"
	c.cars[id].UpdateAItem(w, r)
}

/*DeleteAItem Delete item from the car*/
func (c *Cars) DeleteAItem(w http.ResponseWriter, r *http.Request) {
	id := "1"
	c.cars[id].DeleteAItem(w, r)
}
