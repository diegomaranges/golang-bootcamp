package car

import (
	"fmt"
	"net/http"
)

/*UnitUserActions User interactions with your car*/
type UnitUserActions interface {
}

/*Car Data about diferents cars*/
type Car struct {
}

/*AddItem Add new item in to the car*/
func AddItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*ReturnListOfItems Return list with all elements in the car*/
func ReturnListOfItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*ReturnAItem Return a single item from the car, if not exist return a error*/
func ReturnAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*UpdateAItem Change item from the car*/
func UpdateAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*DeleteAItem Delete item from the car*/
func DeleteAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
