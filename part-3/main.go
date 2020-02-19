package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/db"
)

func main() {
	var keyElement string
	var newElement string
	var keyNewElement string
	var option string
	inputType := bufio.NewScanner(os.Stdin)

	myDataBase := db.CreateNewDBInstance()
	myDataBase.Init()
	myDataBase.LoadFile()

	for {
		myDataBase.PtrintMap()
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
			fmt.Println("Write new element key")
			fmt.Print(">>")
			inputType.Scan()
			keyNewElement = inputType.Text()

			myDataBase.Add(keyNewElement)

		case "retrieve":
			fmt.Println("Write key of element")
			fmt.Print(">>")
			inputType.Scan()
			keyElement = inputType.Text()

			fmt.Println("************************")
			value, err := myDataBase.Retrieve(keyElement)
			if err == nil {
				fmt.Println(value)
			}
			fmt.Println("************************")

		case "update":
			fmt.Println("Write element to replase")
			fmt.Print(">>")
			inputType.Scan()
			keyElement = inputType.Text()

			fmt.Println("Write new value")
			fmt.Print(">>")
			inputType.Scan()
			newElement = inputType.Text()

			myDataBase.Update(keyElement, newElement)

		case "delete":
			fmt.Println("Write element to delete")
			fmt.Print(">>")
			inputType.Scan()
			keyElement = inputType.Text()

			myDataBase.Delete(keyElement)

		case "exit":
			myDataBase.SaveFile()
			return

		default:
			fmt.Println("not valid option")
		}
		fmt.Println("")

	}
}
