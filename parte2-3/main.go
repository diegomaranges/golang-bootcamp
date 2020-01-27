package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var option string
	var stringMap map[string]string
	stringMap = make(map[string]string)
	inputType := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("*************************")
		for key, value := range stringMap {
			fmt.Println(key + ": " + value)
		}
		fmt.Println("*************************")

		fmt.Println("select an option:")
		fmt.Println("add     -> for add a new element")
		fmt.Println("retrive -> for show all elements")
		fmt.Println("update  -> for replace an element")
		fmt.Println("delete  -> for delete an element")
		fmt.Println("exit")
		fmt.Print(">>")

		inputType.Scan()
		option = inputType.Text()
		option = strings.ToLower(option)

		switch option {
		case "add":
			addElement(stringMap)

		case "retrive":
			showElement(stringMap)

		case "update":
			updateElement(stringMap)

		case "delete":
			deleteElement(stringMap)

		case "exit":
			return

		default:
			fmt.Println("not valid option")
			fmt.Println("")
		}

	}
}

func addElement(stringMap map[string]string) {
	var newElement string
	var keyNewElement string
	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write new element key")
	fmt.Print(">>")
	inputType.Scan()
	keyNewElement = inputType.Text()

	_, isUsed := stringMap[keyNewElement]

	if !isUsed {
		fmt.Println("Write new element value")
		fmt.Print(">>")
		inputType.Scan()
		newElement = inputType.Text()
		stringMap[keyNewElement] = newElement
	} else {
		fmt.Println("Key is current used")
	}

	fmt.Println("")
}

func showElement(stringMap map[string]string) {
	var keyElement string

	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write key of element")
	fmt.Print(">>")
	inputType.Scan()
	keyElement = inputType.Text()

	fmt.Println("************************")
	fmt.Println(stringMap[keyElement])
	fmt.Println("************************")
	fmt.Println("")
}

func updateElement(stringMap map[string]string) {
	var newElement string
	var keyNewElement string

	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write element to replase")
	fmt.Print(">>")
	inputType.Scan()
	keyNewElement = inputType.Text()

	_, isUsed := stringMap[keyNewElement]

	if isUsed {
		fmt.Println("Write new value")
		fmt.Print(">>")
		inputType.Scan()
		newElement = inputType.Text()
		stringMap[keyNewElement] = newElement
	} else {
		fmt.Println("Key not found")
	}

	fmt.Println("")
}

func deleteElement(stringMap map[string]string) {
	var elementToDelete string

	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write element to delete")
	fmt.Print(">>")
	inputType.Scan()
	elementToDelete = inputType.Text()

	_, isUsed := stringMap[elementToDelete]

	if isUsed {
		delete(stringMap, elementToDelete)
		fmt.Println("element deleted successful")
	} else {
		fmt.Println("element not found")
	}

	fmt.Println("")
}
