package db

import (
	"errors"
	"sync"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/fileinteraction"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/readapi"
)

/*CRUD All function that you can use in this interface.
Use Init when declare an element with this interface */
type CRUD interface {
	LoadFile() error
	Add(string) error
	ReturnMap() (map[string]*fileinteraction.Items, error)
	Retrieve(string) (fileinteraction.Items, error)
	Update(string, string) error
	Delete(string) error
	SaveFile() error
	DeleteFile() error
}

/*Database Contain a map with the items in the Database */
type Database struct {
	mapInformation map[string]*fileinteraction.Items
	file           *fileinteraction.DestinyFile
	mux            *sync.Mutex
}

/*CreateNewDBInstance create new instance of the object*/
func CreateNewDBInstance(carID string, createNewDB bool) (*Database, error) {
	if createNewDB {
		tempDB := fileinteraction.CreateNewFInstance(carID)
		return nil, tempDB.CreateFile()
	}
	dataBase := &Database{}
	dataBase.mapInformation = make(map[string]*fileinteraction.Items)
	dataBase.file = fileinteraction.CreateNewFInstance(carID)
	dataBase.mux = &sync.Mutex{}
	return dataBase, nil
}

/*LoadFile load file and save information in the db
Pre: Database != nil;
Pos: Return a error if can read the file*/
func (d *Database) LoadFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.file.ReadFile(d.mapInformation)
}

/*Add add item to car
Pre: Database != nil;
Pos: return a error if item does not exist in the API*/
func (d *Database) Add(newID string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	if _, isUsed := d.mapInformation[newID]; isUsed {
		d.mapInformation[newID].Quantity++
		return nil
	}

	result, err := readapi.GetElement(newID)
	if err != nil {
		return errors.New("irem doesnt exist")
	}

	var myNewElement fileinteraction.Items
	myNewElement.Price = result.Price
	myNewElement.Title = result.Title
	myNewElement.Quantity = 1
	d.mapInformation[newID] = &myNewElement

	return nil
}

/*Retrieve return a item from the car
Pre: Database != nil;
Pos: Return a error if item does not exist in the car*/
func (d *Database) Retrieve(id string) (fileinteraction.Items, error) {
	var errorCase fileinteraction.Items

	if d.mapInformation == nil {
		return errorCase, errors.New("map is not initialized")
	}

	d.mux.Lock()
	defer d.mux.Unlock()
	if _, isUsed := d.mapInformation[id]; !isUsed {
		return errorCase, errors.New("item does not exist")
	}

	return *d.mapInformation[id], nil
}

/*ReturnMap return all items from the car
Pre: Database != nil;
Pos: Return the map and nil error*/
func (d *Database) ReturnMap() (map[string]*fileinteraction.Items, error) {
	if d.mapInformation == nil {
		return nil, errors.New("map is not initialized")
	}

	return d.mapInformation, nil
}

/*Update rewrite item from car
Pre: Database != nil;
Pos: return a error if any id item is already used or not exist in the API*/
func (d *Database) Update(actualID string, newID string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	if _, isUsed := d.mapInformation[actualID]; !isUsed {
		return errors.New("Key does not exist")
	}
	if _, isUsed := d.mapInformation[newID]; isUsed {
		return errors.New("Key does not exist")
	}

	result, err := readapi.GetElement(newID)
	if err != nil {
		return errors.New("new item does not exist")
	}
	var myNewElement fileinteraction.Items
	myNewElement.Price = result.Price
	myNewElement.Title = result.Title
	myNewElement.Quantity = d.mapInformation[actualID].Quantity

	d.mapInformation[newID] = &myNewElement
	delete(d.mapInformation, actualID)

	return nil
}

/*Delete remove item from car
Pre: Database != nil;
Pos: Return a error if the car does not have this item*/
func (d *Database) Delete(id string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	if _, isUsed := d.mapInformation[id]; !isUsed {
		return errors.New("Key does not exist")
	}
	delete(d.mapInformation, id)

	return nil
}

/*SaveFile save car
Pre: Database != nil;
Pos: Return a error if can not write the file*/
func (d *Database) SaveFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.file.WriteFile(d.mapInformation)
}

/*DeleteFile delete car
Pre: Database != nil;
Pos: Return a error if does not exist or has any problem with remove file*/
func (d *Database) DeleteFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.file.DeleteFile()
}
