package cars

import "github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/cars/car"

/*Database All functions used for difference between users*/
type Database interface {
}

/*Cars All cars from data base*/
type Cars struct {
	cars map[string]car.Car
}

/*CreateNewInstance Create New Car instance*/
func CreateNewInstance() {

}

/*DeleteInstance Delete Car and alls items included*/
func DeleteInstance() {

}
