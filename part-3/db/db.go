package db

import (
	"errors"
	"fmt"

	"github.corp.globant.com/diego-maranges/golang-bootcamp/part-3/db/fileinteraction"
)

const destinyFile string = "info.txt"

/*CRUD All function that you can use in this interface.
Use Init when declare an element with this interface */
type CRUD interface {
	Init()
	LoadFile() error
	Add(string, string) error
	Retrieve(string) (string, error)
	Update(string, string) error
	Delete(string) error
	SaveFile() error
	PtrintMap() error
}

/*Database Contain a map with the items in the Database */
type Database struct {
	mapInformation map[string]string
	file           *fileinteraction.DestinyFile
}

/*CreateNewDBInstance create new instance of the object*/
func CreateNewDBInstance() *Database {
	return &Database{}
}

/*Init Run first to initilice the Database*/
func (d *Database) Init() {
	d.mapInformation = make(map[string]string)
	d.file = fileinteraction.CreateNewFInstance()
	d.file.SetFile(destinyFile)
}

/*LoadFile load file and save information in the db if have a correct syntax
Pre: Database != nil;
Pos: If have some problem with the file return -1 else, return 0*/
func (d *Database) LoadFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}

	return d.file.ReadFile(d.mapInformation)
}

/*Add add item to db
Pre: Database != nil;
Pos: If key is existent return -1 else, return 0 and add the new item;*/
func (d *Database) Add(keyNewElement string, newElement string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}
	_, isUsed := d.mapInformation[keyNewElement]

	if isUsed {
		return errors.New("Key is already exist")
	}
	d.mapInformation[keyNewElement] = newElement

	return nil
}

/*Retrieve show item from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and the item value;*/
func (d *Database) Retrieve(keyElement string) (string, error) {
	if d.mapInformation == nil {
		return "", errors.New("map is not initialized")
	}

	value, isUsed := d.mapInformation[keyElement]

	if !isUsed {
		return "", errors.New("Key does not exist")
	}

	return value, nil
}

/*Update rewrite item from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and update&return the item value;*/
func (d *Database) Update(keyNewElement string, newElement string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}

	_, isUsed := d.mapInformation[keyNewElement]

	if !isUsed {
		return errors.New("Key does not exist")
	}
	d.mapInformation[keyNewElement] = newElement
	return nil
}

/*Delete remove element from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and remove the item;*/
func (d *Database) Delete(elementToDelete string) error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}

	_, isUsed := d.mapInformation[elementToDelete]
	if !isUsed {
		return errors.New("Key does not exist")
	}
	delete(d.mapInformation, elementToDelete)

	return nil
}

/*SaveFile save db in a file
Pre: Database != nil;
Pos: If have some problem with the file return -1 else, return 0*/
func (d *Database) SaveFile() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}

	return d.file.WriteFile(d.mapInformation)
}

/*PtrintMap show db in the console
Pre: Database != nil;
Pos: Show for console all items in the Database*/
func (d *Database) PtrintMap() error {
	if d.mapInformation == nil {
		return errors.New("map is not initialized")
	}

	fmt.Println("")
	fmt.Println("*************************")
	for key, value := range d.mapInformation {
		fmt.Println(key + ": " + value)
	}
	fmt.Println("*************************")
	return nil
}
