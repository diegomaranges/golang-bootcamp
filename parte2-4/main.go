package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type functionalitys interface {
	add(string, string) int
	retrieve(string) (int, string)
	update(string, string) (int, string)
	delete(string) int
}

type database struct {
	mapInformation map[string]string
	inputType      *bufio.Scanner
}

func (d database) add(keyNewElement string, newElement string) int {
	_, isUsed := d.mapInformation[keyNewElement]

	if !isUsed {
		d.mapInformation[keyNewElement] = newElement

		return 0
	}
	return -1
}

func (d database) retrieve(keyElement string) (int, string) {
	value, isUsed := d.mapInformation[keyElement]

	if !isUsed {
		return -1, ""
	}

	fmt.Println("************************")
	fmt.Println(value)
	fmt.Println("************************")
	return 0, value
}

func (d database) update(keyNewElement string, newElement string) (int, string) {
	_, isUsed := d.mapInformation[keyNewElement]

	if !isUsed {
		fmt.Println("Key not found")

		return -1, ""
	}
	d.mapInformation[keyNewElement] = newElement
	return 0, newElement
}

func (d database) delete(elementToDelete string) int {
	_, isUsed := d.mapInformation[elementToDelete]
	if !isUsed {
		fmt.Println("element not found")

		return -1
	}
	delete(d.mapInformation, elementToDelete)
	fmt.Println("element deleted successful")

	return 0
}

func main() {
	var keyElement string
	var newElement string
	var keyNewElement string
	var option string

	var myDataBase database
	myDataBase.mapInformation = make(map[string]string)
	myDataBase.inputType = bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("*************************")
		for key, value := range myDataBase.mapInformation {
			fmt.Println(key + ": " + value)
		}
		fmt.Println("*************************")

		fmt.Println("select an option:")
		fmt.Println("add     -> for add a new element")
		fmt.Println("retrieve -> for show all elements")
		fmt.Println("update  -> for replace an element")
		fmt.Println("delete  -> for delete an element")
		fmt.Println("exit")
		fmt.Print(">>")

		myDataBase.inputType.Scan()
		option = myDataBase.inputType.Text()
		option = strings.ToLower(option)

		switch option {
		case "add":
			fmt.Println("Write new element key")
			fmt.Print(">>")
			myDataBase.inputType.Scan()
			keyNewElement = myDataBase.inputType.Text()
			fmt.Println("Write new element value")
			fmt.Print(">>")
			myDataBase.inputType.Scan()
			newElement = myDataBase.inputType.Text()

			myDataBase.add(keyNewElement, newElement)

		case "retrieve":
			fmt.Println("Write key of element")
			fmt.Print(">>")
			myDataBase.inputType.Scan()
			keyElement = myDataBase.inputType.Text()

			myDataBase.retrieve(keyElement)

		case "update":
			fmt.Println("Write element to replase")
			fmt.Print(">>")
			myDataBase.inputType.Scan()
			keyElement = myDataBase.inputType.Text()

			fmt.Println("Write new value")
			fmt.Print(">>")
			myDataBase.inputType.Scan()
			newElement = myDataBase.inputType.Text()

			myDataBase.update(keyElement, newElement)

		case "delete":
			fmt.Println("Write element to delete")
			fmt.Print(">>")
			myDataBase.inputType.Scan()
			keyElement = myDataBase.inputType.Text()

			myDataBase.delete(keyElement)

		case "exit":
			return

		default:
			fmt.Println("not valid option")
		}
		fmt.Println("")

	}
}
