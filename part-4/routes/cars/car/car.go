package car

import (
	"fmt"
	"net/http"
)

/*UnitUserActions User interactions with your car*/
type UnitUserActions interface {
	AddItem(w http.ResponseWriter, r *http.Request)
	ReturnListOfItems(w http.ResponseWriter, r *http.Request)
	ReturnAItem(w http.ResponseWriter, r *http.Request)
	UpdateAItem(w http.ResponseWriter, r *http.Request)
	DeleteAItem(w http.ResponseWriter, r *http.Request)
}

/*Item Data for a Item*/
type Item struct {
}

/*Car Data about diferents cars*/
type Car struct {
}

/*CreateNewCar create new empty car*/
func CreateNewCar(w http.ResponseWriter, r *http.Request) *Car {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	return &Car{}
}

/*AddItem Add new item in to the car*/
func (c *Car) AddItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*ReturnListOfItems Return list with all elements in the car*/
func (c *Car) ReturnListOfItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*ReturnAItem Return a single item from the car, if not exist return a error*/
func (c *Car) ReturnAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*UpdateAItem Change item from the car*/
func (c *Car) UpdateAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*DeleteAItem Delete item from the car*/
func (c *Car) DeleteAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
