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
			fmt.Println("")

		case "retrive":
			showElement(stringMap)
			fmt.Println("")

		case "update":
			updateElement(stringMap)
			fmt.Println("")

		case "delete":
			deleteElement(stringMap)
			fmt.Println("")

		case "exit":
			return

		default:
			fmt.Println("not valid option")
			fmt.Println("")
		}

	}
}

func addElement(stringMap map[string]string) int {
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

		return 0
	} else {

		return -1
	}
}

func showElement(stringMap map[string]string) int {
	var keyElement string

	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write key of element")
	fmt.Print(">>")
	inputType.Scan()
	keyElement = inputType.Text()

	value, isUsed := stringMap[keyElement]

	if isUsed {
		fmt.Println("************************")
		fmt.Println(value)
		fmt.Println("************************")

		return 0
	} else {

		return -1
	}
}

func updateElement(stringMap map[string]string) int {
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

		return 0
	} else {
		fmt.Println("Key not found")

		return -1
	}
}

func deleteElement(stringMap map[string]string) int {
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

		return 0
	} else {
		fmt.Println("element not found")

		return -1
	}
}
