package db

import (
	"fmt"
)

/*Functionalitys algo*/
type Functionalitys interface {
	Add(string, string) int
	Retrieve(string) (int, string)
	Update(string, string) (int, string)
	Delete(string) int
	Init()
	PtrintMap()
}

/*Database algo*/
type Database struct {
	mapInformation map[string]string
}

/*Init algo*/
func (d *Database) Init() {
	d.mapInformation = make(map[string]string)
}

/*Add es algo*/
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

/*Retrieve algo*/
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

/*Update algo*/
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

/*Delete algo*/
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

/*PtrintMap algo*/
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
