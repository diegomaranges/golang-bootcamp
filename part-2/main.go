package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-2/db"
)

func main() {
	var keyElement string
	var newElement string
	var keyNewElement string
	var option string
	inputType := bufio.NewScanner(os.Stdin)

	myDataBase := new(db.Database)
	myDataBase.Init()

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
			fmt.Println("Write new element value")
			fmt.Print(">>")
			inputType.Scan()
			newElement = inputType.Text()

			myDataBase.Add(keyNewElement, newElement)

		case "retrieve":
			fmt.Println("Write key of element")
			fmt.Print(">>")
			inputType.Scan()
			keyElement = inputType.Text()

			fmt.Println("************************")
			_, value := myDataBase.Retrieve(keyElement)
			fmt.Println(value)
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
			return

		default:
			fmt.Println("not valid option")
		}
		fmt.Println("")

	}
}
