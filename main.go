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
	var inputS string
	inputType := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("select an option:")
		fmt.Println("add     -> for add a new element")
		fmt.Println("retrive -> for show all elements")
		fmt.Println("update  -> for replace an element")
		fmt.Println("delete  -> for delete an element")
		fmt.Print(">>")

		inputType.Scan()
		option = inputType.Text()
		option = strings.ToLower(option)

		switch option {
		case "add":
			fmt.Println("write the new element")
			fmt.Print(">>")
			inputType.Scan()
			inputS = inputType.Text()
			addElement(inputS)

		case "retrive":
			showElements()

		case "update":
			fmt.Println("update")
			fmt.Println("")

		case "delete":
			inputType.Scan()
			inputS = inputType.Text()
			addElement(inputS)
			fmt.Println("")

		default:
			fmt.Println("not valid option")
			fmt.Println("")
		}

	}
}

func addElement(element string) {
	sliceWithElements = append(sliceWithElements, element)
	fmt.Println("")
}

func deleteElement(element string) {

}

func showElements() {
	fmt.Println("************************")
	for _, element := range sliceWithElements {
		fmt.Println(element)
	}
	fmt.Println("************************")
	fmt.Println("")
}
