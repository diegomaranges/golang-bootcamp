package db

import (
	"fmt"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/db/fileinteraction"
)

const destinyFile string = "info.txt"

/*CRUD All function that you can use in this interface.
Use Init when declare an element with this interface */
type CRUD interface {
	Init()
	LoadFile() int
	Add(string, string) int
	Retrieve(string) (int, string)
	Update(string, string) (int, string)
	Delete(string) int
	SaveFile() int
	PtrintMap()
}

/*Database Contain a map with the items in the Database */
type Database struct {
	mapInformation map[string]string
	file           *fileinteraction.DestinyFile
}

/*Init Run first to initilice the Database*/
func (d *Database) Init() {
	d.mapInformation = make(map[string]string)
	d.file = new(fileinteraction.DestinyFile)
	d.file.SetFile(destinyFile)
}

/*LoadFile load file and save information in the db if have a correct syntax
Pre: Database != nil;
Pos: If have some problem with the file return -1 else, return 0*/
func (d *Database) LoadFile() int {
	if d.mapInformation == nil {
		return -1
	}

	return d.file.ReadFile(d.mapInformation)
}

/*Add add item to db
Pre: Database != nil;
Pos: If key is existent return -1 else, return 0 and add the new item;*/
func (d *Database) Add(keyNewElement string, newElement string) int {
	if d.mapInformation == nil {
		return -1
	}

	_, isUsed := d.mapInformation[keyNewElement]

	if !isUsed {
		d.mapInformation[keyNewElement] = newElement

		return 0
	}
	return -1
}

/*Retrieve show item from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and the item value;*/
func (d *Database) Retrieve(keyElement string) (int, string) {
	if d.mapInformation == nil {
		return -1, ""
	}

	value, isUsed := d.mapInformation[keyElement]

	if !isUsed {
		return -1, ""
	}
	return 0, value
}

/*Update rewrite item from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and update&return the item value;*/
func (d *Database) Update(keyNewElement string, newElement string) (int, string) {
	if d.mapInformation == nil {
		return -1, ""
	}

	_, isUsed := d.mapInformation[keyNewElement]

	if !isUsed {
		fmt.Println("Key not found")

		return -1, ""
	}
	d.mapInformation[keyNewElement] = newElement
	return 0, newElement
}

/*Delete remove element from db
Pre: Database != nil;
Pos: If key is non-existent return -1 else, return 0 and remove the item;*/
func (d *Database) Delete(elementToDelete string) int {
	if d.mapInformation == nil {
		return -1
	}

	_, isUsed := d.mapInformation[elementToDelete]
	if !isUsed {
		fmt.Println("element not found")

		return -1
	}
	delete(d.mapInformation, elementToDelete)

	return 0
}

/*SaveFile save db in a file
Pre: Database != nil;
Pos: If have some problem with the file return -1 else, return 0*/
func (d *Database) SaveFile() int {
	if d.mapInformation == nil {
		return -1
	}

	return d.file.WriteFile(d.mapInformation)
}

/*PtrintMap show db in the console
Pre: Database != nil;
Pos: Show for console all items in the Database*/
func (d *Database) PtrintMap() {
	if d.mapInformation == nil {
		return
	}

	fmt.Println("")
	fmt.Println("*************************")
	for key, value := range d.mapInformation {
		fmt.Println(key + ": " + value)
	}
	fmt.Println("*************************")
}
