package dbcars

import (
	"errors"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/dbcars/dbcar"
)

/*CarsInteractions algo*/
type CarsInteractions interface {
	ReturnListOfCars()
}

/*Cars algo*/
type Cars struct {
	cars map[string]*dbcar.Database
}

/*CreateNewCarsInstance algo*/
func CreateNewCarsInstance() (*Cars, error) {
	firstCar := "1"
	cars := &Cars{}
	cars.cars = make(map[string]*dbcar.Database)
	cars.cars[firstCar] = dbcar.CreateNewDBInstance()
	return cars, nil
}

/*ReturnListOfCars algo*/
func (c *Cars) ReturnListOfCars() (map[string]*dbcar.Database, error) {
	if c.cars == nil {
		return nil, errors.New("map is not initialized")
	}

	return c.cars, nil
}

/*ReturnASpecificCar algo*/
func (c *Cars) ReturnASpecificCar(id string) (*dbcar.Database, error) {
	if c.cars == nil {
		return nil, errors.New("map is not initialized")
	}

	car, isUsed := c.cars[id]
	if !isUsed {
		return nil, errors.New("map is not initialized")
	}
	return car, nil
}

/*CreateNewCar algo*/
func (c *Cars) CreateNewCar(id string) error {
	if c.cars == nil {
		return errors.New("map is not initialized")
	}

	_, isUsed := c.cars[id]
	if isUsed {
		return errors.New("map is not initialized")
	}

	c.cars[id] = dbcar.CreateNewDBInstance()
	return nil
}
