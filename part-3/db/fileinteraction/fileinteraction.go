package fileinteraction

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*FileActions have all function for handle the interecctions with files*/
type FileActions interface {
	SetFile(string)
	ReadFile(map[string]string) error
	WriteFile(map[string]string) error
	ReturnDestinyFile() string
}

/*DestinyFile save Destiny File for all information as a string*/
type DestinyFile struct {
	destinyFile string
}

/*CreateNewFInstance create new instance of the object*/
func CreateNewFInstance() *DestinyFile {
	return &DestinyFile{}
}

/*SetFile save new route file*/
func (d *DestinyFile) SetFile(destiny string) {
	d.destinyFile = destiny
}

/*ReadFile reed information from file saved
Warning: If the map element have some information it will delete
Pre: externalMap different to nil*/
func (d *DestinyFile) ReadFile(externalMap map[string]string) error {
	if externalMap == nil {
		return errors.New("map is not initialized")
	}

	file, err := os.OpenFile(d.destinyFile, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	defer file.Close()

	if err != nil {
		return err
	}

	fileBytes := make([]byte, 100)
	_, err = file.ReadAt(fileBytes, 0)

	if err != io.EOF && err != nil {
		return err
	}

	externalMap = make(map[string]string)

	sliceWithElements := strings.Split(string(fileBytes), "\n")

	for _, row := range sliceWithElements {
		keyAndValue := strings.Split(row, " ")

		if len(keyAndValue) == 2 {
			externalMap[keyAndValue[0]] = keyAndValue[1]
		}
	}

	return nil
}

/*WriteFile write information in file saved
Pre: externalMap different to nil*/
func (d *DestinyFile) WriteFile(externalMap map[string]string) error {
	if externalMap == nil {
		return errors.New("map is not initialized")
	}

	first := true
	file, err := os.OpenFile(d.destinyFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	if err != nil {
		return err
	}

	for key, value := range externalMap {
		element := []byte(key + " " + value + "\n")
		if first {
			err = ioutil.WriteFile(d.destinyFile, element, 077)
			first = false
			if err != nil {
				return err
			}

		} else {
			_, err = file.WriteString(string(element))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*ReturnDestinyFile return a route string saved*/
func (d *DestinyFile) ReturnDestinyFile() string {
	return d.destinyFile
}
