package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func UploadFile(file string) map[string]string {
	stringMap := make(map[string]string)

	elements, err := ioutil.ReadFile(file)
	folder, err := os.OpenFile(file, os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	folder.Close()

	if err == nil {
		sliceWithElements := strings.Split(string(elements), "\n")
		for _, row := range sliceWithElements {
			keyAndValue := strings.Split(row, " ")
			fmt.Println(row)
			if len(keyAndValue) == 2 {
				stringMap[keyAndValue[0]] = keyAndValue[1]
			}
		}
	} else {
		panic(err)
	}

	return stringMap
}

func WriteFolder(file string, stringMap map[string]string) {
	var first bool = true
	folder, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err == nil {
		for key, value := range stringMap {
			element := []byte(key + " " + value + "\n")
			if first {
				er := ioutil.WriteFile(file, element, 077)
				first = false
				if er != nil {
					fmt.Println(er)
				}
			} else {
				_, er := folder.WriteString(string(element))
				if er != nil {
					fmt.Println(er)
				}
			}
		}
		folder.Close()
	} else {
		fmt.Println(err)
	}
}

/*
http://man7.org/linux/man-pages/man2/openat.2.html
func main() {
	new()
	text, err := ioutil.ReadFile("info.txt")
	if err == nil {
		fmt.Println(string(text))
	} else {
		panic(err)
	}
}

func scan() {
	text, err := ioutil.ReadFile("info.txt")
}

func write() {
	newText := []byte(os.Args[1])
	//writeFile := ioutil.WriteFile("info.txt", newText, 077)
	writeFile := ioutil.WriteFile("info.txt", newText, 077)
}

func new() {
	//algo := "asdacas"
	folder, err := os.OpenFile("info.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	fmt.Println(folder)
	fmt.Println(err)
	variable, er := folder.WriteString("algo")
	fmt.Println(variable)
	fmt.Println(er)
	folder.Close()
}*/
