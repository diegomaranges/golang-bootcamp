package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sliceWithElements []string

func main() {
	var option string
	inputType := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("select an option:")
		fmt.Println("add     -> for add a new element")
		fmt.Println("retrieve -> for show all elements")
		fmt.Println("update  -> for replace an element")
		fmt.Println("delete  -> for delete an element")
		fmt.Println("exit")
		fmt.Print(">>")

		inputType.Scan()
		option = inputType.Text()
		option = strings.ToLower(option)

		switch option {
		case "add":
			addElement()

		case "retrieve":
			showElements()

		case "update":
			updateElement()

		case "delete":
			deleteElement()

		case "exit":
			return

		default:
			fmt.Println("not valid option")
			fmt.Println("")
		}

	}
}

func addElement() {
	var newElement string
	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write the new element")
	fmt.Print(">>")
	inputType.Scan()
	newElement = inputType.Text()

	sliceWithElements = append(sliceWithElements, newElement)
	fmt.Println("")
}

func updateElement() {
	var currentElement string
	var newElement string
	var notFound bool = true

	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write element to replase")
	fmt.Print(">>")
	inputType.Scan()
	currentElement = inputType.Text()
	fmt.Println(currentElement)

	for i := 0; i < len(sliceWithElements) && notFound; i++ {
		if currentElement == sliceWithElements[i] {
			fmt.Println("Write the new element")
			fmt.Print(">>")
			inputType.Scan()
			newElement = inputType.Text()

			sliceWithElements[i] = newElement
			notFound = false
		}
	}

	if notFound {
		fmt.Println("Element not found")
	}

	fmt.Println("")
}

func deleteElement() {
	var currentElement string
	var notFound bool = true

	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("Write element to delete")
	fmt.Print(">>")
	inputType.Scan()
	currentElement = inputType.Text()
	fmt.Println(currentElement)

	for i := 0; i < len(sliceWithElements) && notFound; i++ {
		if currentElement == sliceWithElements[i] {
			temp := sliceWithElements[0:i]
			temp = append(temp, sliceWithElements[(i+1):len(sliceWithElements)]...)
			sliceWithElements = temp
			notFound = false
			fmt.Println("Element deleted")
		}
	}

	if notFound {
		fmt.Println("Element not found")
	}

	fmt.Println("")

}

func showElements() {
	fmt.Println("************************")

	for _, element := range sliceWithElements {
		fmt.Println(element)
	}

	fmt.Println("************************")
	fmt.Println("")
}
