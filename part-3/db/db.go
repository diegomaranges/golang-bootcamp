package db

import (
	"errors"
	"fmt"
	"sync"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/db/fileinteraction"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/db/readapi"
)

const destinyFile string = "db"

/*CRUD All function that you can use in this interface.
Use Init when declare an element with this interface */
type CRUD interface {
	LoadFile() error
	Add(string) error
	Retrieve(string) (fileinteraction.Items, error)
	Update(string, string) error
	Delete(string) error
	SaveFile() error
	PtrintMap() error
}

/*Database Contain a map with the items in the Database */
type Database struct {
	mapInformation map[string]*fileinteraction.Items
	file           *fileinteraction.DestinyFile
	mux            *sync.Mutex
}

/*CreateNewDBInstance create new instance of the object*/
func CreateNewDBInstance() *Database {
	dataBase := &Database{}
	dataBase.mapInformation = make(map[string]*fileinteraction.Items)
	dataBase.file = fileinteraction.CreateNewFInstance()
	dataBase.file.SetFile(destinyFile)
	dataBase.mux = &sync.Mutex{}
	return dataBase
}

/*LoadFile load file and save information in the db if have a correct syntax
Pre: Database != nil;
Pos: If have some problem with the file return -1 else, return 0*/
func (d *Database) LoadFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.file.ReadFile(d.mapInformation)
}

/*Add add item to db
Pre: Database != nil;
Pos: If key is existent return -1 else, return 0 and add the new item;*/
func (d *Database) Add(newID string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	_, isUsed := d.mapInformation[newID]

	if isUsed {
		d.mapInformation[newID].Quantity++
		return nil
	}
	result, err := readapi.GetElement(newID)
	if err != nil {
		return errors.New("element doesnt exist")
	}

	var myNewElement fileinteraction.Items
	myNewElement.Price = result.Price
	myNewElement.Title = result.Title
	myNewElement.Quantity = 1
	d.mapInformation[newID] = &myNewElement

	return nil
}

/*Retrieve show item from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and the item value;*/
func (d *Database) Retrieve(id string) (fileinteraction.Items, error) {
	var errorCase fileinteraction.Items

	if d.mapInformation == nil {
		return errorCase, errors.New("map is not initialized")
	}

	d.mux.Lock()
	defer d.mux.Unlock()
	value, isUsed := d.mapInformation[id]

	if !isUsed {

		return errorCase, errors.New("element does not exist")
	}

	return *value, nil
}

/*Update rewrite item from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and update&return the item value;*/
func (d *Database) Update(actualID string, newID string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	_, isUsed := d.mapInformation[actualID]
	if !isUsed {
		return errors.New("Key does not exist")
	}
	_, isUsed = d.mapInformation[newID]
	if isUsed {
		return errors.New("Key does not exist")
	}

	result, err := readapi.GetElement(newID)
	fmt.Println(err)
	if err != nil {
		return errors.New("new element does not exist")
	}

	var myNewElement fileinteraction.Items
	myNewElement.Price = result.Price
	myNewElement.Title = result.Title
	myNewElement.Quantity = d.mapInformation[actualID].Quantity

	d.mapInformation[newID] = &myNewElement
	delete(d.mapInformation, actualID)

	return nil
}

/*Delete remove element from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and remove the item;*/
func (d *Database) Delete(id string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	_, isUsed := d.mapInformation[id]
	if !isUsed {
		return errors.New("Key does not exist")
	}
	delete(d.mapInformation, id)

	return nil
}

/*SaveFile save db in a file
Pre: Database != nil;
Pos: If have some problem with the file return -1 else, return 0*/
func (d *Database) SaveFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	return d.file.WriteFile(d.mapInformation)
}

/*PtrintMap show db in the console
Pre: Database != nil;
Pos: Show for console all items in the Database*/
func (d *Database) PtrintMap() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	d.mux.Lock()
	defer d.mux.Unlock()

	fmt.Println("")
	fmt.Println("*************************")
	for k, v := range d.mapInformation {
		fmt.Print(k + ": ")
		fmt.Println(*v)
	}
	fmt.Println("*************************")
	return nil
}
