package fileinteraction

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*FileActions have all function for handle the interecctions with files*/
type FileActions interface {
	SetFile(string)
	ReadFile(map[string]string) int
	WriteFile(map[string]string) int
	ReturnDestinyFile() string
}

/*DestinyFile save Destiny File for all information as a string*/
type DestinyFile struct {
	destinyFile string
}

/*SetFile save new route file*/
func (d *DestinyFile) SetFile(destiny string) {
	d.destinyFile = destiny
}

/*ReadFile reed information from file saved
Pre: externalMap different to nil*/
func (d *DestinyFile) ReadFile(externalMap map[string]string) int {
	if externalMap == nil {
		return -1
	}

	file, err := os.OpenFile(d.destinyFile, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	fileBytes := make([]byte, 100)
	_, err = file.ReadAt(fileBytes, 0)

	if err != io.EOF && err != nil {
		panic(err)
	}

	sliceWithElements := strings.Split(string(fileBytes), "\n")

	for _, row := range sliceWithElements {
		keyAndValue := strings.Split(row, " ")
		fmt.Println(row)
		if len(keyAndValue) == 2 {
			externalMap[keyAndValue[0]] = keyAndValue[1]
		}
	}

	return 0
}

/*WriteFile write information in file saved
Pre: externalMap different to nil*/
func (d *DestinyFile) WriteFile(externalMap map[string]string) int {
	if externalMap == nil {
		return -1
	}

	first := true
	file, err := os.OpenFile(d.destinyFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	if err == nil {
		for key, value := range externalMap {
			element := []byte(key + " " + value + "\n")
			if first {
				er := ioutil.WriteFile(d.destinyFile, element, 077)
				first = false
				if er != nil {
					fmt.Println(er)
					break
				}

			} else {
				_, er := file.WriteString(string(element))
				if er != nil {
					fmt.Println(er)
					break
				}
			}
		}
		return 0
	}
	fmt.Println(err)
	return -1
}

/*ReturnDestinyFile return a route string saved*/
func (d *DestinyFile) ReturnDestinyFile() string {
	return d.destinyFile
}
