package db

import (
	"errors"
	"sync"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/mongodb"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/mongodb/fileinteraction"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/readapi"
)

/*CRUD All function that you can use in this interface.*/
type CRUD interface {
	Add(string) error
	Retrieve(string) (fileinteraction.Item, error)
	RetrieveAll() (fileinteraction.Items, error)
	Update(string, string) error
	Delete(string) error
	DeleteCar() error
}

/*Database Contain a map with the items in the Database */
type Database struct {
	mux *sync.Mutex
	db  *mongodb.MongoStruct
}

/*CreateNewDBInstance create new instance of the object using the directory and carID

If createNewDB is true CreateNewDBInstance only create the file and return a nil *Database*/
func CreateNewDBInstance(directory string, carID string, newDB bool) (*Database, error) {
	var err error
	dataBase := &Database{}
	dataBase.mux = &sync.Mutex{}
	dataBase.db, err = mongodb.CreateNewDBInstance(directory, carID, newDB)
	return dataBase, err
}

/*Add add item to car

Pre: Database != nil;

Pos: return a error if item does not exist in the API*/
func (d *Database) Add(id string) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	item, err := d.db.ReturnItem(id)
	if err == nil {
		return d.db.AddItem(item)
	}

	result, err := readapi.GetElement(id)
	if err != nil {
		return errors.New("item does not exist")
	}

	var newElement fileinteraction.Item
	newElement.ID = result.ID
	newElement.Price = result.Price
	newElement.Title = result.Title
	newElement.Quantity = 1

	return d.db.AddItem(newElement)
}

/*Retrieve return a item from the car

Pre: Database != nil;

Pos: Return a error if item does not exist in the car*/
func (d *Database) Retrieve(id string) (fileinteraction.Item, error) {
	d.mux.Lock()
	defer d.mux.Unlock()
	return d.db.ReturnItem(id)
}

/*RetrieveAll return all items from the car

Pre: Database != nil;

Pos: Return the map and nil error*/
func (d *Database) RetrieveAll() (fileinteraction.Items, error) {
	return d.db.ReturnAllItems()
}

/*Update rewrite item from car

Pre: Database != nil;

Pos: return a error if any id item is already used or not exist in the API*/
func (d *Database) Update(id string, quantity int) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.db.UpdateItem(id, quantity)
}

/*Delete remove item from car

Pre: Database != nil;

Pos: Return a error if the car does not have this item*/
func (d *Database) Delete(id string) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.db.DeleteItem(id)
}

/*DeleteCar algo*/
func (d *Database) DeleteCar() error {
	return d.db.DeleteColection()
}
